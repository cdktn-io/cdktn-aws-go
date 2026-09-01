package awseventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPipesPipe_SourceParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ActivemqBrokerParameters() AwsPipesPipe_ActivemqBrokerParametersPropertyOutputReference
	// Experimental.
	ActivemqBrokerParametersInput() *AwsPipesPipe_ActivemqBrokerParametersProperty
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
	DynamodbStreamParameters() AwsPipesPipe_DynamodbStreamParametersPropertyOutputReference
	// Experimental.
	DynamodbStreamParametersInput() *AwsPipesPipe_DynamodbStreamParametersProperty
	// Experimental.
	FilterCriteria() AwsPipesPipe_FilterCriteriaPropertyOutputReference
	// Experimental.
	FilterCriteriaInput() *AwsPipesPipe_FilterCriteriaProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsPipesPipe_SourceParametersProperty
	// Experimental.
	SetInternalValue(val *AwsPipesPipe_SourceParametersProperty)
	// Experimental.
	KinesisStreamParameters() AwsPipesPipe_SourceParametersKinesisStreamParametersPropertyOutputReference
	// Experimental.
	KinesisStreamParametersInput() *AwsPipesPipe_SourceParametersKinesisStreamParametersProperty
	// Experimental.
	ManagedStreamingKafkaParameters() AwsPipesPipe_ManagedStreamingKafkaParametersPropertyOutputReference
	// Experimental.
	ManagedStreamingKafkaParametersInput() *AwsPipesPipe_ManagedStreamingKafkaParametersProperty
	// Experimental.
	RabbitmqBrokerParameters() AwsPipesPipe_RabbitmqBrokerParametersPropertyOutputReference
	// Experimental.
	RabbitmqBrokerParametersInput() *AwsPipesPipe_RabbitmqBrokerParametersProperty
	// Experimental.
	SelfManagedKafkaParameters() AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference
	// Experimental.
	SelfManagedKafkaParametersInput() *AwsPipesPipe_SelfManagedKafkaParametersProperty
	// Experimental.
	SqsQueueParameters() AwsPipesPipe_SourceParametersSqsQueueParametersPropertyOutputReference
	// Experimental.
	SqsQueueParametersInput() *AwsPipesPipe_SourceParametersSqsQueueParametersProperty
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
	PutActivemqBrokerParameters(value *AwsPipesPipe_ActivemqBrokerParametersProperty)
	// Experimental.
	PutDynamodbStreamParameters(value *AwsPipesPipe_DynamodbStreamParametersProperty)
	// Experimental.
	PutFilterCriteria(value *AwsPipesPipe_FilterCriteriaProperty)
	// Experimental.
	PutKinesisStreamParameters(value *AwsPipesPipe_SourceParametersKinesisStreamParametersProperty)
	// Experimental.
	PutManagedStreamingKafkaParameters(value *AwsPipesPipe_ManagedStreamingKafkaParametersProperty)
	// Experimental.
	PutRabbitmqBrokerParameters(value *AwsPipesPipe_RabbitmqBrokerParametersProperty)
	// Experimental.
	PutSelfManagedKafkaParameters(value *AwsPipesPipe_SelfManagedKafkaParametersProperty)
	// Experimental.
	PutSqsQueueParameters(value *AwsPipesPipe_SourceParametersSqsQueueParametersProperty)
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

// The jsii proxy struct for AwsPipesPipe_SourceParametersPropertyOutputReference
type jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) ActivemqBrokerParameters() AwsPipesPipe_ActivemqBrokerParametersPropertyOutputReference {
	var returns AwsPipesPipe_ActivemqBrokerParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"activemqBrokerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) ActivemqBrokerParametersInput() *AwsPipesPipe_ActivemqBrokerParametersProperty {
	var returns *AwsPipesPipe_ActivemqBrokerParametersProperty
	_jsii_.Get(
		j,
		"activemqBrokerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) DynamodbStreamParameters() AwsPipesPipe_DynamodbStreamParametersPropertyOutputReference {
	var returns AwsPipesPipe_DynamodbStreamParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"dynamodbStreamParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) DynamodbStreamParametersInput() *AwsPipesPipe_DynamodbStreamParametersProperty {
	var returns *AwsPipesPipe_DynamodbStreamParametersProperty
	_jsii_.Get(
		j,
		"dynamodbStreamParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) FilterCriteria() AwsPipesPipe_FilterCriteriaPropertyOutputReference {
	var returns AwsPipesPipe_FilterCriteriaPropertyOutputReference
	_jsii_.Get(
		j,
		"filterCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) FilterCriteriaInput() *AwsPipesPipe_FilterCriteriaProperty {
	var returns *AwsPipesPipe_FilterCriteriaProperty
	_jsii_.Get(
		j,
		"filterCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) InternalValue() *AwsPipesPipe_SourceParametersProperty {
	var returns *AwsPipesPipe_SourceParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) KinesisStreamParameters() AwsPipesPipe_SourceParametersKinesisStreamParametersPropertyOutputReference {
	var returns AwsPipesPipe_SourceParametersKinesisStreamParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) KinesisStreamParametersInput() *AwsPipesPipe_SourceParametersKinesisStreamParametersProperty {
	var returns *AwsPipesPipe_SourceParametersKinesisStreamParametersProperty
	_jsii_.Get(
		j,
		"kinesisStreamParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) ManagedStreamingKafkaParameters() AwsPipesPipe_ManagedStreamingKafkaParametersPropertyOutputReference {
	var returns AwsPipesPipe_ManagedStreamingKafkaParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"managedStreamingKafkaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) ManagedStreamingKafkaParametersInput() *AwsPipesPipe_ManagedStreamingKafkaParametersProperty {
	var returns *AwsPipesPipe_ManagedStreamingKafkaParametersProperty
	_jsii_.Get(
		j,
		"managedStreamingKafkaParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) RabbitmqBrokerParameters() AwsPipesPipe_RabbitmqBrokerParametersPropertyOutputReference {
	var returns AwsPipesPipe_RabbitmqBrokerParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"rabbitmqBrokerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) RabbitmqBrokerParametersInput() *AwsPipesPipe_RabbitmqBrokerParametersProperty {
	var returns *AwsPipesPipe_RabbitmqBrokerParametersProperty
	_jsii_.Get(
		j,
		"rabbitmqBrokerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) SelfManagedKafkaParameters() AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference {
	var returns AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"selfManagedKafkaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) SelfManagedKafkaParametersInput() *AwsPipesPipe_SelfManagedKafkaParametersProperty {
	var returns *AwsPipesPipe_SelfManagedKafkaParametersProperty
	_jsii_.Get(
		j,
		"selfManagedKafkaParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) SqsQueueParameters() AwsPipesPipe_SourceParametersSqsQueueParametersPropertyOutputReference {
	var returns AwsPipesPipe_SourceParametersSqsQueueParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"sqsQueueParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) SqsQueueParametersInput() *AwsPipesPipe_SourceParametersSqsQueueParametersProperty {
	var returns *AwsPipesPipe_SourceParametersSqsQueueParametersProperty
	_jsii_.Get(
		j,
		"sqsQueueParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPipesPipe_SourceParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPipesPipe_SourceParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPipesPipe_SourceParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipesPipe.SourceParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPipesPipe_SourceParametersPropertyOutputReference_Override(a AwsPipesPipe_SourceParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipesPipe.SourceParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference)SetInternalValue(val *AwsPipesPipe_SourceParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) PutActivemqBrokerParameters(value *AwsPipesPipe_ActivemqBrokerParametersProperty) {
	if err := a.validatePutActivemqBrokerParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putActivemqBrokerParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) PutDynamodbStreamParameters(value *AwsPipesPipe_DynamodbStreamParametersProperty) {
	if err := a.validatePutDynamodbStreamParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynamodbStreamParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) PutFilterCriteria(value *AwsPipesPipe_FilterCriteriaProperty) {
	if err := a.validatePutFilterCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilterCriteria",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) PutKinesisStreamParameters(value *AwsPipesPipe_SourceParametersKinesisStreamParametersProperty) {
	if err := a.validatePutKinesisStreamParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisStreamParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) PutManagedStreamingKafkaParameters(value *AwsPipesPipe_ManagedStreamingKafkaParametersProperty) {
	if err := a.validatePutManagedStreamingKafkaParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedStreamingKafkaParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) PutRabbitmqBrokerParameters(value *AwsPipesPipe_RabbitmqBrokerParametersProperty) {
	if err := a.validatePutRabbitmqBrokerParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRabbitmqBrokerParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) PutSelfManagedKafkaParameters(value *AwsPipesPipe_SelfManagedKafkaParametersProperty) {
	if err := a.validatePutSelfManagedKafkaParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSelfManagedKafkaParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) PutSqsQueueParameters(value *AwsPipesPipe_SourceParametersSqsQueueParametersProperty) {
	if err := a.validatePutSqsQueueParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqsQueueParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) ResetActivemqBrokerParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetActivemqBrokerParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) ResetDynamodbStreamParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamodbStreamParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) ResetFilterCriteria() {
	_jsii_.InvokeVoid(
		a,
		"resetFilterCriteria",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) ResetKinesisStreamParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisStreamParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) ResetManagedStreamingKafkaParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedStreamingKafkaParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) ResetRabbitmqBrokerParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetRabbitmqBrokerParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) ResetSelfManagedKafkaParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetSelfManagedKafkaParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) ResetSqsQueueParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetSqsQueueParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPipesPipe_SourceParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

