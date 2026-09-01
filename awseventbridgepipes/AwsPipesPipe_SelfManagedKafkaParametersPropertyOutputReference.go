package awseventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdditionalBootstrapServers() *[]*string
	// Experimental.
	SetAdditionalBootstrapServers(val *[]*string)
	// Experimental.
	AdditionalBootstrapServersInput() *[]*string
	// Experimental.
	BatchSize() *float64
	// Experimental.
	SetBatchSize(val *float64)
	// Experimental.
	BatchSizeInput() *float64
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
	ConsumerGroupId() *string
	// Experimental.
	SetConsumerGroupId(val *string)
	// Experimental.
	ConsumerGroupIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Credentials() AwsPipesPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference
	// Experimental.
	CredentialsInput() *AwsPipesPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsPipesPipe_SelfManagedKafkaParametersProperty
	// Experimental.
	SetInternalValue(val *AwsPipesPipe_SelfManagedKafkaParametersProperty)
	// Experimental.
	MaximumBatchingWindowInSeconds() *float64
	// Experimental.
	SetMaximumBatchingWindowInSeconds(val *float64)
	// Experimental.
	MaximumBatchingWindowInSecondsInput() *float64
	// Experimental.
	ServerRootCaCertificate() *string
	// Experimental.
	SetServerRootCaCertificate(val *string)
	// Experimental.
	ServerRootCaCertificateInput() *string
	// Experimental.
	StartingPosition() *string
	// Experimental.
	SetStartingPosition(val *string)
	// Experimental.
	StartingPositionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TopicName() *string
	// Experimental.
	SetTopicName(val *string)
	// Experimental.
	TopicNameInput() *string
	// Experimental.
	Vpc() AwsPipesPipe_VpcPropertyOutputReference
	// Experimental.
	VpcInput() *AwsPipesPipe_VpcProperty
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
	PutCredentials(value *AwsPipesPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty)
	// Experimental.
	PutVpc(value *AwsPipesPipe_VpcProperty)
	// Experimental.
	ResetAdditionalBootstrapServers()
	// Experimental.
	ResetBatchSize()
	// Experimental.
	ResetConsumerGroupId()
	// Experimental.
	ResetCredentials()
	// Experimental.
	ResetMaximumBatchingWindowInSeconds()
	// Experimental.
	ResetServerRootCaCertificate()
	// Experimental.
	ResetStartingPosition()
	// Experimental.
	ResetVpc()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference
type jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) AdditionalBootstrapServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalBootstrapServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) AdditionalBootstrapServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalBootstrapServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) BatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) BatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) ConsumerGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) ConsumerGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) Credentials() AwsPipesPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference {
	var returns AwsPipesPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) CredentialsInput() *AwsPipesPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty {
	var returns *AwsPipesPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) InternalValue() *AwsPipesPipe_SelfManagedKafkaParametersProperty {
	var returns *AwsPipesPipe_SelfManagedKafkaParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) MaximumBatchingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) MaximumBatchingWindowInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) ServerRootCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverRootCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) ServerRootCaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverRootCaCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) StartingPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) StartingPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) TopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) TopicNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) Vpc() AwsPipesPipe_VpcPropertyOutputReference {
	var returns AwsPipesPipe_VpcPropertyOutputReference
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) VpcInput() *AwsPipesPipe_VpcProperty {
	var returns *AwsPipesPipe_VpcProperty
	_jsii_.Get(
		j,
		"vpcInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipesPipe.SelfManagedKafkaParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference_Override(a AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipesPipe.SelfManagedKafkaParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference)SetAdditionalBootstrapServers(val *[]*string) {
	if err := j.validateSetAdditionalBootstrapServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalBootstrapServers",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference)SetBatchSize(val *float64) {
	if err := j.validateSetBatchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference)SetConsumerGroupId(val *string) {
	if err := j.validateSetConsumerGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerGroupId",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference)SetInternalValue(val *AwsPipesPipe_SelfManagedKafkaParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference)SetMaximumBatchingWindowInSeconds(val *float64) {
	if err := j.validateSetMaximumBatchingWindowInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBatchingWindowInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference)SetServerRootCaCertificate(val *string) {
	if err := j.validateSetServerRootCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverRootCaCertificate",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference)SetStartingPosition(val *string) {
	if err := j.validateSetStartingPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingPosition",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference)SetTopicName(val *string) {
	if err := j.validateSetTopicNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicName",
		val,
	)
}

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) PutCredentials(value *AwsPipesPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty) {
	if err := a.validatePutCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCredentials",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) PutVpc(value *AwsPipesPipe_VpcProperty) {
	if err := a.validatePutVpcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpc",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) ResetAdditionalBootstrapServers() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalBootstrapServers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) ResetBatchSize() {
	_jsii_.InvokeVoid(
		a,
		"resetBatchSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) ResetConsumerGroupId() {
	_jsii_.InvokeVoid(
		a,
		"resetConsumerGroupId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) ResetCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) ResetMaximumBatchingWindowInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumBatchingWindowInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) ResetServerRootCaCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetServerRootCaCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) ResetStartingPosition() {
	_jsii_.InvokeVoid(
		a,
		"resetStartingPosition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) ResetVpc() {
	_jsii_.InvokeVoid(
		a,
		"resetVpc",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPipesPipe_SelfManagedKafkaParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

