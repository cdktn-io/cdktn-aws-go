package awseventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPipe_SourceParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ActivemqBrokerParameters() TfPipe_ActivemqBrokerParametersPropertyOutputReference
	// Experimental.
	ActivemqBrokerParametersInput() *TfPipe_ActivemqBrokerParametersProperty
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
	DynamodbStreamParameters() TfPipe_DynamodbStreamParametersPropertyOutputReference
	// Experimental.
	DynamodbStreamParametersInput() *TfPipe_DynamodbStreamParametersProperty
	// Experimental.
	FilterCriteria() TfPipe_FilterCriteriaPropertyOutputReference
	// Experimental.
	FilterCriteriaInput() *TfPipe_FilterCriteriaProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfPipe_SourceParametersProperty
	// Experimental.
	SetInternalValue(val *TfPipe_SourceParametersProperty)
	// Experimental.
	KinesisStreamParameters() TfPipe_SourceParametersKinesisStreamParametersPropertyOutputReference
	// Experimental.
	KinesisStreamParametersInput() *TfPipe_SourceParametersKinesisStreamParametersProperty
	// Experimental.
	ManagedStreamingKafkaParameters() TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference
	// Experimental.
	ManagedStreamingKafkaParametersInput() *TfPipe_ManagedStreamingKafkaParametersProperty
	// Experimental.
	RabbitmqBrokerParameters() TfPipe_RabbitmqBrokerParametersPropertyOutputReference
	// Experimental.
	RabbitmqBrokerParametersInput() *TfPipe_RabbitmqBrokerParametersProperty
	// Experimental.
	SelfManagedKafkaParameters() TfPipe_SelfManagedKafkaParametersPropertyOutputReference
	// Experimental.
	SelfManagedKafkaParametersInput() *TfPipe_SelfManagedKafkaParametersProperty
	// Experimental.
	SqsQueueParameters() TfPipe_SourceParametersSqsQueueParametersPropertyOutputReference
	// Experimental.
	SqsQueueParametersInput() *TfPipe_SourceParametersSqsQueueParametersProperty
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
	PutActivemqBrokerParameters(value *TfPipe_ActivemqBrokerParametersProperty)
	// Experimental.
	PutDynamodbStreamParameters(value *TfPipe_DynamodbStreamParametersProperty)
	// Experimental.
	PutFilterCriteria(value *TfPipe_FilterCriteriaProperty)
	// Experimental.
	PutKinesisStreamParameters(value *TfPipe_SourceParametersKinesisStreamParametersProperty)
	// Experimental.
	PutManagedStreamingKafkaParameters(value *TfPipe_ManagedStreamingKafkaParametersProperty)
	// Experimental.
	PutRabbitmqBrokerParameters(value *TfPipe_RabbitmqBrokerParametersProperty)
	// Experimental.
	PutSelfManagedKafkaParameters(value *TfPipe_SelfManagedKafkaParametersProperty)
	// Experimental.
	PutSqsQueueParameters(value *TfPipe_SourceParametersSqsQueueParametersProperty)
	// Experimental.
	ResetActivemqBrokerParameters()
	// Experimental.
	ResetDynamodbStreamParameters()
	// Experimental.
	ResetFilterCriteria()
	// Experimental.
	ResetKinesisStreamParameters()
	// Experimental.
	ResetManagedStreamingKafkaParameters()
	// Experimental.
	ResetRabbitmqBrokerParameters()
	// Experimental.
	ResetSelfManagedKafkaParameters()
	// Experimental.
	ResetSqsQueueParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPipe_SourceParametersPropertyOutputReference
type jsiiProxy_TfPipe_SourceParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) ActivemqBrokerParameters() TfPipe_ActivemqBrokerParametersPropertyOutputReference {
	var returns TfPipe_ActivemqBrokerParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"activemqBrokerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) ActivemqBrokerParametersInput() *TfPipe_ActivemqBrokerParametersProperty {
	var returns *TfPipe_ActivemqBrokerParametersProperty
	_jsii_.Get(
		j,
		"activemqBrokerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) DynamodbStreamParameters() TfPipe_DynamodbStreamParametersPropertyOutputReference {
	var returns TfPipe_DynamodbStreamParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"dynamodbStreamParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) DynamodbStreamParametersInput() *TfPipe_DynamodbStreamParametersProperty {
	var returns *TfPipe_DynamodbStreamParametersProperty
	_jsii_.Get(
		j,
		"dynamodbStreamParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) FilterCriteria() TfPipe_FilterCriteriaPropertyOutputReference {
	var returns TfPipe_FilterCriteriaPropertyOutputReference
	_jsii_.Get(
		j,
		"filterCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) FilterCriteriaInput() *TfPipe_FilterCriteriaProperty {
	var returns *TfPipe_FilterCriteriaProperty
	_jsii_.Get(
		j,
		"filterCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) InternalValue() *TfPipe_SourceParametersProperty {
	var returns *TfPipe_SourceParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) KinesisStreamParameters() TfPipe_SourceParametersKinesisStreamParametersPropertyOutputReference {
	var returns TfPipe_SourceParametersKinesisStreamParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) KinesisStreamParametersInput() *TfPipe_SourceParametersKinesisStreamParametersProperty {
	var returns *TfPipe_SourceParametersKinesisStreamParametersProperty
	_jsii_.Get(
		j,
		"kinesisStreamParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) ManagedStreamingKafkaParameters() TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference {
	var returns TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"managedStreamingKafkaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) ManagedStreamingKafkaParametersInput() *TfPipe_ManagedStreamingKafkaParametersProperty {
	var returns *TfPipe_ManagedStreamingKafkaParametersProperty
	_jsii_.Get(
		j,
		"managedStreamingKafkaParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) RabbitmqBrokerParameters() TfPipe_RabbitmqBrokerParametersPropertyOutputReference {
	var returns TfPipe_RabbitmqBrokerParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"rabbitmqBrokerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) RabbitmqBrokerParametersInput() *TfPipe_RabbitmqBrokerParametersProperty {
	var returns *TfPipe_RabbitmqBrokerParametersProperty
	_jsii_.Get(
		j,
		"rabbitmqBrokerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) SelfManagedKafkaParameters() TfPipe_SelfManagedKafkaParametersPropertyOutputReference {
	var returns TfPipe_SelfManagedKafkaParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"selfManagedKafkaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) SelfManagedKafkaParametersInput() *TfPipe_SelfManagedKafkaParametersProperty {
	var returns *TfPipe_SelfManagedKafkaParametersProperty
	_jsii_.Get(
		j,
		"selfManagedKafkaParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) SqsQueueParameters() TfPipe_SourceParametersSqsQueueParametersPropertyOutputReference {
	var returns TfPipe_SourceParametersSqsQueueParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"sqsQueueParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) SqsQueueParametersInput() *TfPipe_SourceParametersSqsQueueParametersProperty {
	var returns *TfPipe_SourceParametersSqsQueueParametersProperty
	_jsii_.Get(
		j,
		"sqsQueueParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPipe_SourceParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPipe_SourceParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPipe_SourceParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPipe_SourceParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.SourceParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPipe_SourceParametersPropertyOutputReference_Override(t TfPipe_SourceParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.SourceParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference)SetInternalValue(val *TfPipe_SourceParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) PutActivemqBrokerParameters(value *TfPipe_ActivemqBrokerParametersProperty) {
	if err := t.validatePutActivemqBrokerParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putActivemqBrokerParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) PutDynamodbStreamParameters(value *TfPipe_DynamodbStreamParametersProperty) {
	if err := t.validatePutDynamodbStreamParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDynamodbStreamParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) PutFilterCriteria(value *TfPipe_FilterCriteriaProperty) {
	if err := t.validatePutFilterCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFilterCriteria",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) PutKinesisStreamParameters(value *TfPipe_SourceParametersKinesisStreamParametersProperty) {
	if err := t.validatePutKinesisStreamParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesisStreamParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) PutManagedStreamingKafkaParameters(value *TfPipe_ManagedStreamingKafkaParametersProperty) {
	if err := t.validatePutManagedStreamingKafkaParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putManagedStreamingKafkaParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) PutRabbitmqBrokerParameters(value *TfPipe_RabbitmqBrokerParametersProperty) {
	if err := t.validatePutRabbitmqBrokerParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRabbitmqBrokerParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) PutSelfManagedKafkaParameters(value *TfPipe_SelfManagedKafkaParametersProperty) {
	if err := t.validatePutSelfManagedKafkaParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSelfManagedKafkaParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) PutSqsQueueParameters(value *TfPipe_SourceParametersSqsQueueParametersProperty) {
	if err := t.validatePutSqsQueueParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSqsQueueParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) ResetActivemqBrokerParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetActivemqBrokerParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) ResetDynamodbStreamParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetDynamodbStreamParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) ResetFilterCriteria() {
	_jsii_.InvokeVoid(
		t,
		"resetFilterCriteria",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) ResetKinesisStreamParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesisStreamParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) ResetManagedStreamingKafkaParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedStreamingKafkaParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) ResetRabbitmqBrokerParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetRabbitmqBrokerParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) ResetSelfManagedKafkaParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetSelfManagedKafkaParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) ResetSqsQueueParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetSqsQueueParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPipe_SourceParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

