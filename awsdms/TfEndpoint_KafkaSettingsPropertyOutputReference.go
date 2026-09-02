package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpoint_KafkaSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Broker() *string
	// Experimental.
	SetBroker(val *string)
	// Experimental.
	BrokerInput() *string
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
	Fqn() *string
	// Experimental.
	IncludeControlDetails() interface{}
	// Experimental.
	SetIncludeControlDetails(val interface{})
	// Experimental.
	IncludeControlDetailsInput() interface{}
	// Experimental.
	IncludeNullAndEmpty() interface{}
	// Experimental.
	SetIncludeNullAndEmpty(val interface{})
	// Experimental.
	IncludeNullAndEmptyInput() interface{}
	// Experimental.
	IncludePartitionValue() interface{}
	// Experimental.
	SetIncludePartitionValue(val interface{})
	// Experimental.
	IncludePartitionValueInput() interface{}
	// Experimental.
	IncludeTableAlterOperations() interface{}
	// Experimental.
	SetIncludeTableAlterOperations(val interface{})
	// Experimental.
	IncludeTableAlterOperationsInput() interface{}
	// Experimental.
	IncludeTransactionDetails() interface{}
	// Experimental.
	SetIncludeTransactionDetails(val interface{})
	// Experimental.
	IncludeTransactionDetailsInput() interface{}
	// Experimental.
	InternalValue() *TfEndpoint_KafkaSettingsProperty
	// Experimental.
	SetInternalValue(val *TfEndpoint_KafkaSettingsProperty)
	// Experimental.
	MessageFormat() *string
	// Experimental.
	SetMessageFormat(val *string)
	// Experimental.
	MessageFormatInput() *string
	// Experimental.
	MessageMaxBytes() *float64
	// Experimental.
	SetMessageMaxBytes(val *float64)
	// Experimental.
	MessageMaxBytesInput() *float64
	// Experimental.
	NoHexPrefix() interface{}
	// Experimental.
	SetNoHexPrefix(val interface{})
	// Experimental.
	NoHexPrefixInput() interface{}
	// Experimental.
	PartitionIncludeSchemaTable() interface{}
	// Experimental.
	SetPartitionIncludeSchemaTable(val interface{})
	// Experimental.
	PartitionIncludeSchemaTableInput() interface{}
	// Experimental.
	SaslMechanism() *string
	// Experimental.
	SetSaslMechanism(val *string)
	// Experimental.
	SaslMechanismInput() *string
	// Experimental.
	SaslPassword() *string
	// Experimental.
	SetSaslPassword(val *string)
	// Experimental.
	SaslPasswordInput() *string
	// Experimental.
	SaslUsername() *string
	// Experimental.
	SetSaslUsername(val *string)
	// Experimental.
	SaslUsernameInput() *string
	// Experimental.
	SecurityProtocol() *string
	// Experimental.
	SetSecurityProtocol(val *string)
	// Experimental.
	SecurityProtocolInput() *string
	// Experimental.
	SslCaCertificateArn() *string
	// Experimental.
	SetSslCaCertificateArn(val *string)
	// Experimental.
	SslCaCertificateArnInput() *string
	// Experimental.
	SslClientCertificateArn() *string
	// Experimental.
	SetSslClientCertificateArn(val *string)
	// Experimental.
	SslClientCertificateArnInput() *string
	// Experimental.
	SslClientKeyArn() *string
	// Experimental.
	SetSslClientKeyArn(val *string)
	// Experimental.
	SslClientKeyArnInput() *string
	// Experimental.
	SslClientKeyPassword() *string
	// Experimental.
	SetSslClientKeyPassword(val *string)
	// Experimental.
	SslClientKeyPasswordInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Topic() *string
	// Experimental.
	SetTopic(val *string)
	// Experimental.
	TopicInput() *string
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
	ResetIncludeControlDetails()
	// Experimental.
	ResetIncludeNullAndEmpty()
	// Experimental.
	ResetIncludePartitionValue()
	// Experimental.
	ResetIncludeTableAlterOperations()
	// Experimental.
	ResetIncludeTransactionDetails()
	// Experimental.
	ResetMessageFormat()
	// Experimental.
	ResetMessageMaxBytes()
	// Experimental.
	ResetNoHexPrefix()
	// Experimental.
	ResetPartitionIncludeSchemaTable()
	// Experimental.
	ResetSaslMechanism()
	// Experimental.
	ResetSaslPassword()
	// Experimental.
	ResetSaslUsername()
	// Experimental.
	ResetSecurityProtocol()
	// Experimental.
	ResetSslCaCertificateArn()
	// Experimental.
	ResetSslClientCertificateArn()
	// Experimental.
	ResetSslClientKeyArn()
	// Experimental.
	ResetSslClientKeyPassword()
	// Experimental.
	ResetTopic()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEndpoint_KafkaSettingsPropertyOutputReference
type jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) Broker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"broker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) BrokerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brokerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) IncludeControlDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeControlDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) IncludeControlDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeControlDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) IncludeNullAndEmpty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNullAndEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) IncludeNullAndEmptyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNullAndEmptyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) IncludePartitionValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includePartitionValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) IncludePartitionValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includePartitionValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) IncludeTableAlterOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTableAlterOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) IncludeTableAlterOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTableAlterOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) IncludeTransactionDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTransactionDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) IncludeTransactionDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTransactionDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) InternalValue() *TfEndpoint_KafkaSettingsProperty {
	var returns *TfEndpoint_KafkaSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) MessageFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) MessageFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) MessageMaxBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageMaxBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) MessageMaxBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageMaxBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) NoHexPrefix() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHexPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) NoHexPrefixInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHexPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) PartitionIncludeSchemaTable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionIncludeSchemaTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) PartitionIncludeSchemaTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionIncludeSchemaTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) SaslMechanism() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslMechanism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) SaslMechanismInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslMechanismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) SaslPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) SaslPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) SaslUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) SaslUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) SecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) SecurityProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) SslCaCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) SslCaCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) SslClientCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) SslClientCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) SslClientKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) SslClientKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) SslClientKeyPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) SslClientKeyPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) Topic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) TopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEndpoint_KafkaSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEndpoint_KafkaSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEndpoint_KafkaSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.TfEndpoint.KafkaSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEndpoint_KafkaSettingsPropertyOutputReference_Override(t TfEndpoint_KafkaSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.TfEndpoint.KafkaSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetBroker(val *string) {
	if err := j.validateSetBrokerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"broker",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetIncludeControlDetails(val interface{}) {
	if err := j.validateSetIncludeControlDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeControlDetails",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetIncludeNullAndEmpty(val interface{}) {
	if err := j.validateSetIncludeNullAndEmptyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeNullAndEmpty",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetIncludePartitionValue(val interface{}) {
	if err := j.validateSetIncludePartitionValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePartitionValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetIncludeTableAlterOperations(val interface{}) {
	if err := j.validateSetIncludeTableAlterOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTableAlterOperations",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetIncludeTransactionDetails(val interface{}) {
	if err := j.validateSetIncludeTransactionDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTransactionDetails",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetInternalValue(val *TfEndpoint_KafkaSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetMessageFormat(val *string) {
	if err := j.validateSetMessageFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageFormat",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetMessageMaxBytes(val *float64) {
	if err := j.validateSetMessageMaxBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageMaxBytes",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetNoHexPrefix(val interface{}) {
	if err := j.validateSetNoHexPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noHexPrefix",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetPartitionIncludeSchemaTable(val interface{}) {
	if err := j.validateSetPartitionIncludeSchemaTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionIncludeSchemaTable",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetSaslMechanism(val *string) {
	if err := j.validateSetSaslMechanismParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saslMechanism",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetSaslPassword(val *string) {
	if err := j.validateSetSaslPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saslPassword",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetSaslUsername(val *string) {
	if err := j.validateSetSaslUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saslUsername",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetSecurityProtocol(val *string) {
	if err := j.validateSetSecurityProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProtocol",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetSslCaCertificateArn(val *string) {
	if err := j.validateSetSslCaCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCaCertificateArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetSslClientCertificateArn(val *string) {
	if err := j.validateSetSslClientCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslClientCertificateArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetSslClientKeyArn(val *string) {
	if err := j.validateSetSslClientKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslClientKeyArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetSslClientKeyPassword(val *string) {
	if err := j.validateSetSslClientKeyPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslClientKeyPassword",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference)SetTopic(val *string) {
	if err := j.validateSetTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topic",
		val,
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetIncludeControlDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeControlDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetIncludeNullAndEmpty() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeNullAndEmpty",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetIncludePartitionValue() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludePartitionValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetIncludeTableAlterOperations() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeTableAlterOperations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetIncludeTransactionDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeTransactionDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetMessageFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetMessageFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetMessageMaxBytes() {
	_jsii_.InvokeVoid(
		t,
		"resetMessageMaxBytes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetNoHexPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetNoHexPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetPartitionIncludeSchemaTable() {
	_jsii_.InvokeVoid(
		t,
		"resetPartitionIncludeSchemaTable",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetSaslMechanism() {
	_jsii_.InvokeVoid(
		t,
		"resetSaslMechanism",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetSaslPassword() {
	_jsii_.InvokeVoid(
		t,
		"resetSaslPassword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetSaslUsername() {
	_jsii_.InvokeVoid(
		t,
		"resetSaslUsername",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetSecurityProtocol() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityProtocol",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetSslCaCertificateArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSslCaCertificateArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetSslClientCertificateArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSslClientCertificateArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetSslClientKeyArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSslClientKeyArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetSslClientKeyPassword() {
	_jsii_.InvokeVoid(
		t,
		"resetSslClientKeyPassword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ResetTopic() {
	_jsii_.InvokeVoid(
		t,
		"resetTopic",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEndpoint_KafkaSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

