package dms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/dms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/dms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEndpoint_KafkaSettingsPropertyOutputReference interface {
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
	InternalValue() *AwsEndpoint_KafkaSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsEndpoint_KafkaSettingsProperty)
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

// The jsii proxy struct for AwsEndpoint_KafkaSettingsPropertyOutputReference
type jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) Broker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"broker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) BrokerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brokerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) IncludeControlDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeControlDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) IncludeControlDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeControlDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) IncludeNullAndEmpty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNullAndEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) IncludeNullAndEmptyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNullAndEmptyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) IncludePartitionValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includePartitionValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) IncludePartitionValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includePartitionValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) IncludeTableAlterOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTableAlterOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) IncludeTableAlterOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTableAlterOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) IncludeTransactionDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTransactionDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) IncludeTransactionDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTransactionDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) InternalValue() *AwsEndpoint_KafkaSettingsProperty {
	var returns *AwsEndpoint_KafkaSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) MessageFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) MessageFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) MessageMaxBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageMaxBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) MessageMaxBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageMaxBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) NoHexPrefix() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHexPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) NoHexPrefixInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHexPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) PartitionIncludeSchemaTable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionIncludeSchemaTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) PartitionIncludeSchemaTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionIncludeSchemaTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) SaslMechanism() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslMechanism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) SaslMechanismInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslMechanismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) SaslPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) SaslPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) SaslUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) SaslUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) SecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) SecurityProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) SslCaCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) SslCaCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) SslClientCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) SslClientCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) SslClientKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) SslClientKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) SslClientKeyPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) SslClientKeyPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) Topic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) TopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEndpoint_KafkaSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEndpoint_KafkaSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEndpoint_KafkaSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.AwsEndpoint.KafkaSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEndpoint_KafkaSettingsPropertyOutputReference_Override(a AwsEndpoint_KafkaSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.AwsEndpoint.KafkaSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetBroker(val *string) {
	if err := j.validateSetBrokerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"broker",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetIncludeControlDetails(val interface{}) {
	if err := j.validateSetIncludeControlDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeControlDetails",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetIncludeNullAndEmpty(val interface{}) {
	if err := j.validateSetIncludeNullAndEmptyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeNullAndEmpty",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetIncludePartitionValue(val interface{}) {
	if err := j.validateSetIncludePartitionValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePartitionValue",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetIncludeTableAlterOperations(val interface{}) {
	if err := j.validateSetIncludeTableAlterOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTableAlterOperations",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetIncludeTransactionDetails(val interface{}) {
	if err := j.validateSetIncludeTransactionDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTransactionDetails",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetInternalValue(val *AwsEndpoint_KafkaSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetMessageFormat(val *string) {
	if err := j.validateSetMessageFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageFormat",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetMessageMaxBytes(val *float64) {
	if err := j.validateSetMessageMaxBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageMaxBytes",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetNoHexPrefix(val interface{}) {
	if err := j.validateSetNoHexPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noHexPrefix",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetPartitionIncludeSchemaTable(val interface{}) {
	if err := j.validateSetPartitionIncludeSchemaTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionIncludeSchemaTable",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetSaslMechanism(val *string) {
	if err := j.validateSetSaslMechanismParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saslMechanism",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetSaslPassword(val *string) {
	if err := j.validateSetSaslPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saslPassword",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetSaslUsername(val *string) {
	if err := j.validateSetSaslUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saslUsername",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetSecurityProtocol(val *string) {
	if err := j.validateSetSecurityProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProtocol",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetSslCaCertificateArn(val *string) {
	if err := j.validateSetSslCaCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCaCertificateArn",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetSslClientCertificateArn(val *string) {
	if err := j.validateSetSslClientCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslClientCertificateArn",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetSslClientKeyArn(val *string) {
	if err := j.validateSetSslClientKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslClientKeyArn",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetSslClientKeyPassword(val *string) {
	if err := j.validateSetSslClientKeyPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslClientKeyPassword",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference)SetTopic(val *string) {
	if err := j.validateSetTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topic",
		val,
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetIncludeControlDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeControlDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetIncludeNullAndEmpty() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeNullAndEmpty",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetIncludePartitionValue() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludePartitionValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetIncludeTableAlterOperations() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeTableAlterOperations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetIncludeTransactionDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeTransactionDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetMessageFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetMessageFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetMessageMaxBytes() {
	_jsii_.InvokeVoid(
		a,
		"resetMessageMaxBytes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetNoHexPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetNoHexPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetPartitionIncludeSchemaTable() {
	_jsii_.InvokeVoid(
		a,
		"resetPartitionIncludeSchemaTable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetSaslMechanism() {
	_jsii_.InvokeVoid(
		a,
		"resetSaslMechanism",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetSaslPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetSaslPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetSaslUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetSaslUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetSecurityProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetSslCaCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSslCaCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetSslClientCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSslClientCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetSslClientKeyArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSslClientKeyArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetSslClientKeyPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetSslClientKeyPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ResetTopic() {
	_jsii_.InvokeVoid(
		a,
		"resetTopic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEndpoint_KafkaSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

