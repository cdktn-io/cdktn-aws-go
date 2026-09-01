package awseventbridgescheduler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgescheduler/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgescheduler/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSchedulerSchedule_TargetPropertyOutputReference interface {
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
	DeadLetterConfig() AwsSchedulerSchedule_DeadLetterConfigPropertyOutputReference
	// Experimental.
	DeadLetterConfigInput() *AwsSchedulerSchedule_DeadLetterConfigProperty
	// Experimental.
	EcsParameters() AwsSchedulerSchedule_EcsParametersPropertyOutputReference
	// Experimental.
	EcsParametersInput() *AwsSchedulerSchedule_EcsParametersProperty
	// Experimental.
	EventbridgeParameters() AwsSchedulerSchedule_EventbridgeParametersPropertyOutputReference
	// Experimental.
	EventbridgeParametersInput() *AwsSchedulerSchedule_EventbridgeParametersProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Input() *string
	// Experimental.
	SetInput(val *string)
	// Experimental.
	InputInput() *string
	// Experimental.
	InternalValue() *AwsSchedulerSchedule_TargetProperty
	// Experimental.
	SetInternalValue(val *AwsSchedulerSchedule_TargetProperty)
	// Experimental.
	KinesisParameters() AwsSchedulerSchedule_KinesisParametersPropertyOutputReference
	// Experimental.
	KinesisParametersInput() *AwsSchedulerSchedule_KinesisParametersProperty
	// Experimental.
	RetryPolicy() AwsSchedulerSchedule_RetryPolicyPropertyOutputReference
	// Experimental.
	RetryPolicyInput() *AwsSchedulerSchedule_RetryPolicyProperty
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	SagemakerPipelineParameters() AwsSchedulerSchedule_SagemakerPipelineParametersPropertyOutputReference
	// Experimental.
	SagemakerPipelineParametersInput() *AwsSchedulerSchedule_SagemakerPipelineParametersProperty
	// Experimental.
	SqsParameters() AwsSchedulerSchedule_SqsParametersPropertyOutputReference
	// Experimental.
	SqsParametersInput() *AwsSchedulerSchedule_SqsParametersProperty
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
	PutDeadLetterConfig(value *AwsSchedulerSchedule_DeadLetterConfigProperty)
	// Experimental.
	PutEcsParameters(value *AwsSchedulerSchedule_EcsParametersProperty)
	// Experimental.
	PutEventbridgeParameters(value *AwsSchedulerSchedule_EventbridgeParametersProperty)
	// Experimental.
	PutKinesisParameters(value *AwsSchedulerSchedule_KinesisParametersProperty)
	// Experimental.
	PutRetryPolicy(value *AwsSchedulerSchedule_RetryPolicyProperty)
	// Experimental.
	PutSagemakerPipelineParameters(value *AwsSchedulerSchedule_SagemakerPipelineParametersProperty)
	// Experimental.
	PutSqsParameters(value *AwsSchedulerSchedule_SqsParametersProperty)
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

// The jsii proxy struct for AwsSchedulerSchedule_TargetPropertyOutputReference
type jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) DeadLetterConfig() AwsSchedulerSchedule_DeadLetterConfigPropertyOutputReference {
	var returns AwsSchedulerSchedule_DeadLetterConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) DeadLetterConfigInput() *AwsSchedulerSchedule_DeadLetterConfigProperty {
	var returns *AwsSchedulerSchedule_DeadLetterConfigProperty
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) EcsParameters() AwsSchedulerSchedule_EcsParametersPropertyOutputReference {
	var returns AwsSchedulerSchedule_EcsParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"ecsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) EcsParametersInput() *AwsSchedulerSchedule_EcsParametersProperty {
	var returns *AwsSchedulerSchedule_EcsParametersProperty
	_jsii_.Get(
		j,
		"ecsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) EventbridgeParameters() AwsSchedulerSchedule_EventbridgeParametersPropertyOutputReference {
	var returns AwsSchedulerSchedule_EventbridgeParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"eventbridgeParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) EventbridgeParametersInput() *AwsSchedulerSchedule_EventbridgeParametersProperty {
	var returns *AwsSchedulerSchedule_EventbridgeParametersProperty
	_jsii_.Get(
		j,
		"eventbridgeParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) InputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) InternalValue() *AwsSchedulerSchedule_TargetProperty {
	var returns *AwsSchedulerSchedule_TargetProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) KinesisParameters() AwsSchedulerSchedule_KinesisParametersPropertyOutputReference {
	var returns AwsSchedulerSchedule_KinesisParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) KinesisParametersInput() *AwsSchedulerSchedule_KinesisParametersProperty {
	var returns *AwsSchedulerSchedule_KinesisParametersProperty
	_jsii_.Get(
		j,
		"kinesisParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) RetryPolicy() AwsSchedulerSchedule_RetryPolicyPropertyOutputReference {
	var returns AwsSchedulerSchedule_RetryPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) RetryPolicyInput() *AwsSchedulerSchedule_RetryPolicyProperty {
	var returns *AwsSchedulerSchedule_RetryPolicyProperty
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) SagemakerPipelineParameters() AwsSchedulerSchedule_SagemakerPipelineParametersPropertyOutputReference {
	var returns AwsSchedulerSchedule_SagemakerPipelineParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"sagemakerPipelineParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) SagemakerPipelineParametersInput() *AwsSchedulerSchedule_SagemakerPipelineParametersProperty {
	var returns *AwsSchedulerSchedule_SagemakerPipelineParametersProperty
	_jsii_.Get(
		j,
		"sagemakerPipelineParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) SqsParameters() AwsSchedulerSchedule_SqsParametersPropertyOutputReference {
	var returns AwsSchedulerSchedule_SqsParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"sqsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) SqsParametersInput() *AwsSchedulerSchedule_SqsParametersProperty {
	var returns *AwsSchedulerSchedule_SqsParametersProperty
	_jsii_.Get(
		j,
		"sqsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSchedulerSchedule_TargetPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSchedulerSchedule_TargetPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSchedulerSchedule_TargetPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-scheduler.AwsSchedulerSchedule.TargetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSchedulerSchedule_TargetPropertyOutputReference_Override(a AwsSchedulerSchedule_TargetPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-scheduler.AwsSchedulerSchedule.TargetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference)SetInput(val *string) {
	if err := j.validateSetInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference)SetInternalValue(val *AwsSchedulerSchedule_TargetProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) PutDeadLetterConfig(value *AwsSchedulerSchedule_DeadLetterConfigProperty) {
	if err := a.validatePutDeadLetterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) PutEcsParameters(value *AwsSchedulerSchedule_EcsParametersProperty) {
	if err := a.validatePutEcsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcsParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) PutEventbridgeParameters(value *AwsSchedulerSchedule_EventbridgeParametersProperty) {
	if err := a.validatePutEventbridgeParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEventbridgeParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) PutKinesisParameters(value *AwsSchedulerSchedule_KinesisParametersProperty) {
	if err := a.validatePutKinesisParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) PutRetryPolicy(value *AwsSchedulerSchedule_RetryPolicyProperty) {
	if err := a.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) PutSagemakerPipelineParameters(value *AwsSchedulerSchedule_SagemakerPipelineParametersProperty) {
	if err := a.validatePutSagemakerPipelineParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSagemakerPipelineParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) PutSqsParameters(value *AwsSchedulerSchedule_SqsParametersProperty) {
	if err := a.validatePutSqsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqsParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) ResetEcsParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetEcsParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) ResetEventbridgeParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetEventbridgeParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) ResetInput() {
	_jsii_.InvokeVoid(
		a,
		"resetInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) ResetKinesisParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) ResetSagemakerPipelineParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetSagemakerPipelineParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) ResetSqsParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetSqsParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSchedulerSchedule_TargetPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

