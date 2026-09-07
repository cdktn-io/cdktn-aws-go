package eventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/eventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/eventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPipe_SourceParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ActivemqBrokerParameters() AwsPipe_ActivemqBrokerParametersPropertyOutputReference
	// Experimental.
	ActivemqBrokerParametersInput() *AwsPipe_ActivemqBrokerParametersProperty
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
	DynamodbStreamParameters() AwsPipe_DynamodbStreamParametersPropertyOutputReference
	// Experimental.
	DynamodbStreamParametersInput() *AwsPipe_DynamodbStreamParametersProperty
	// Experimental.
	FilterCriteria() AwsPipe_FilterCriteriaPropertyOutputReference
	// Experimental.
	FilterCriteriaInput() *AwsPipe_FilterCriteriaProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsPipe_SourceParametersProperty
	// Experimental.
	SetInternalValue(val *AwsPipe_SourceParametersProperty)
	// Experimental.
	KinesisStreamParameters() AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference
	// Experimental.
	KinesisStreamParametersInput() *AwsPipe_SourceParametersKinesisStreamParametersProperty
	// Experimental.
	ManagedStreamingKafkaParameters() AwsPipe_ManagedStreamingKafkaParametersPropertyOutputReference
	// Experimental.
	ManagedStreamingKafkaParametersInput() *AwsPipe_ManagedStreamingKafkaParametersProperty
	// Experimental.
	RabbitmqBrokerParameters() AwsPipe_RabbitmqBrokerParametersPropertyOutputReference
	// Experimental.
	RabbitmqBrokerParametersInput() *AwsPipe_RabbitmqBrokerParametersProperty
	// Experimental.
	SelfManagedKafkaParameters() AwsPipe_SelfManagedKafkaParametersPropertyOutputReference
	// Experimental.
	SelfManagedKafkaParametersInput() *AwsPipe_SelfManagedKafkaParametersProperty
	// Experimental.
	SqsQueueParameters() AwsPipe_SourceParametersSqsQueueParametersPropertyOutputReference
	// Experimental.
	SqsQueueParametersInput() *AwsPipe_SourceParametersSqsQueueParametersProperty
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
	PutActivemqBrokerParameters(value *AwsPipe_ActivemqBrokerParametersProperty)
	// Experimental.
	PutDynamodbStreamParameters(value *AwsPipe_DynamodbStreamParametersProperty)
	// Experimental.
	PutFilterCriteria(value *AwsPipe_FilterCriteriaProperty)
	// Experimental.
	PutKinesisStreamParameters(value *AwsPipe_SourceParametersKinesisStreamParametersProperty)
	// Experimental.
	PutManagedStreamingKafkaParameters(value *AwsPipe_ManagedStreamingKafkaParametersProperty)
	// Experimental.
	PutRabbitmqBrokerParameters(value *AwsPipe_RabbitmqBrokerParametersProperty)
	// Experimental.
	PutSelfManagedKafkaParameters(value *AwsPipe_SelfManagedKafkaParametersProperty)
	// Experimental.
	PutSqsQueueParameters(value *AwsPipe_SourceParametersSqsQueueParametersProperty)
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

// The jsii proxy struct for AwsPipe_SourceParametersPropertyOutputReference
type jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) ActivemqBrokerParameters() AwsPipe_ActivemqBrokerParametersPropertyOutputReference {
	var returns AwsPipe_ActivemqBrokerParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"activemqBrokerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) ActivemqBrokerParametersInput() *AwsPipe_ActivemqBrokerParametersProperty {
	var returns *AwsPipe_ActivemqBrokerParametersProperty
	_jsii_.Get(
		j,
		"activemqBrokerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) DynamodbStreamParameters() AwsPipe_DynamodbStreamParametersPropertyOutputReference {
	var returns AwsPipe_DynamodbStreamParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"dynamodbStreamParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) DynamodbStreamParametersInput() *AwsPipe_DynamodbStreamParametersProperty {
	var returns *AwsPipe_DynamodbStreamParametersProperty
	_jsii_.Get(
		j,
		"dynamodbStreamParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) FilterCriteria() AwsPipe_FilterCriteriaPropertyOutputReference {
	var returns AwsPipe_FilterCriteriaPropertyOutputReference
	_jsii_.Get(
		j,
		"filterCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) FilterCriteriaInput() *AwsPipe_FilterCriteriaProperty {
	var returns *AwsPipe_FilterCriteriaProperty
	_jsii_.Get(
		j,
		"filterCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) InternalValue() *AwsPipe_SourceParametersProperty {
	var returns *AwsPipe_SourceParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) KinesisStreamParameters() AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference {
	var returns AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) KinesisStreamParametersInput() *AwsPipe_SourceParametersKinesisStreamParametersProperty {
	var returns *AwsPipe_SourceParametersKinesisStreamParametersProperty
	_jsii_.Get(
		j,
		"kinesisStreamParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) ManagedStreamingKafkaParameters() AwsPipe_ManagedStreamingKafkaParametersPropertyOutputReference {
	var returns AwsPipe_ManagedStreamingKafkaParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"managedStreamingKafkaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) ManagedStreamingKafkaParametersInput() *AwsPipe_ManagedStreamingKafkaParametersProperty {
	var returns *AwsPipe_ManagedStreamingKafkaParametersProperty
	_jsii_.Get(
		j,
		"managedStreamingKafkaParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) RabbitmqBrokerParameters() AwsPipe_RabbitmqBrokerParametersPropertyOutputReference {
	var returns AwsPipe_RabbitmqBrokerParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"rabbitmqBrokerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) RabbitmqBrokerParametersInput() *AwsPipe_RabbitmqBrokerParametersProperty {
	var returns *AwsPipe_RabbitmqBrokerParametersProperty
	_jsii_.Get(
		j,
		"rabbitmqBrokerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) SelfManagedKafkaParameters() AwsPipe_SelfManagedKafkaParametersPropertyOutputReference {
	var returns AwsPipe_SelfManagedKafkaParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"selfManagedKafkaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) SelfManagedKafkaParametersInput() *AwsPipe_SelfManagedKafkaParametersProperty {
	var returns *AwsPipe_SelfManagedKafkaParametersProperty
	_jsii_.Get(
		j,
		"selfManagedKafkaParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) SqsQueueParameters() AwsPipe_SourceParametersSqsQueueParametersPropertyOutputReference {
	var returns AwsPipe_SourceParametersSqsQueueParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"sqsQueueParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) SqsQueueParametersInput() *AwsPipe_SourceParametersSqsQueueParametersProperty {
	var returns *AwsPipe_SourceParametersSqsQueueParametersProperty
	_jsii_.Get(
		j,
		"sqsQueueParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPipe_SourceParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPipe_SourceParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPipe_SourceParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipe.SourceParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPipe_SourceParametersPropertyOutputReference_Override(a AwsPipe_SourceParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipe.SourceParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference)SetInternalValue(val *AwsPipe_SourceParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) PutActivemqBrokerParameters(value *AwsPipe_ActivemqBrokerParametersProperty) {
	if err := a.validatePutActivemqBrokerParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putActivemqBrokerParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) PutDynamodbStreamParameters(value *AwsPipe_DynamodbStreamParametersProperty) {
	if err := a.validatePutDynamodbStreamParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynamodbStreamParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) PutFilterCriteria(value *AwsPipe_FilterCriteriaProperty) {
	if err := a.validatePutFilterCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilterCriteria",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) PutKinesisStreamParameters(value *AwsPipe_SourceParametersKinesisStreamParametersProperty) {
	if err := a.validatePutKinesisStreamParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisStreamParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) PutManagedStreamingKafkaParameters(value *AwsPipe_ManagedStreamingKafkaParametersProperty) {
	if err := a.validatePutManagedStreamingKafkaParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedStreamingKafkaParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) PutRabbitmqBrokerParameters(value *AwsPipe_RabbitmqBrokerParametersProperty) {
	if err := a.validatePutRabbitmqBrokerParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRabbitmqBrokerParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) PutSelfManagedKafkaParameters(value *AwsPipe_SelfManagedKafkaParametersProperty) {
	if err := a.validatePutSelfManagedKafkaParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSelfManagedKafkaParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) PutSqsQueueParameters(value *AwsPipe_SourceParametersSqsQueueParametersProperty) {
	if err := a.validatePutSqsQueueParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqsQueueParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) ResetActivemqBrokerParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetActivemqBrokerParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) ResetDynamodbStreamParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamodbStreamParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) ResetFilterCriteria() {
	_jsii_.InvokeVoid(
		a,
		"resetFilterCriteria",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) ResetKinesisStreamParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisStreamParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) ResetManagedStreamingKafkaParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedStreamingKafkaParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) ResetRabbitmqBrokerParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetRabbitmqBrokerParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) ResetSelfManagedKafkaParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetSelfManagedKafkaParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) ResetSqsQueueParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetSqsQueueParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPipe_SourceParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

