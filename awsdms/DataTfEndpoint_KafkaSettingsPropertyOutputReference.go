package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfEndpoint_KafkaSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Broker() *string
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
	IncludeControlDetails() cdktn.IResolvable
	// Experimental.
	IncludeNullAndEmpty() cdktn.IResolvable
	// Experimental.
	IncludePartitionValue() cdktn.IResolvable
	// Experimental.
	IncludeTableAlterOperations() cdktn.IResolvable
	// Experimental.
	IncludeTransactionDetails() cdktn.IResolvable
	// Experimental.
	InternalValue() *DataTfEndpoint_KafkaSettingsProperty
	// Experimental.
	SetInternalValue(val *DataTfEndpoint_KafkaSettingsProperty)
	// Experimental.
	MessageFormat() *string
	// Experimental.
	MessageMaxBytes() *float64
	// Experimental.
	NoHexPrefix() cdktn.IResolvable
	// Experimental.
	PartitionIncludeSchemaTable() cdktn.IResolvable
	// Experimental.
	SaslMechanism() *string
	// Experimental.
	SaslPassword() *string
	// Experimental.
	SaslUsername() *string
	// Experimental.
	SecurityProtocol() *string
	// Experimental.
	SslCaCertificateArn() *string
	// Experimental.
	SslClientCertificateArn() *string
	// Experimental.
	SslClientKeyArn() *string
	// Experimental.
	SslClientKeyPassword() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataTfEndpoint_KafkaSettingsPropertyOutputReference
type jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) Broker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"broker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) IncludeControlDetails() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"includeControlDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) IncludeNullAndEmpty() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"includeNullAndEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) IncludePartitionValue() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"includePartitionValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) IncludeTableAlterOperations() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"includeTableAlterOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) IncludeTransactionDetails() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"includeTransactionDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) InternalValue() *DataTfEndpoint_KafkaSettingsProperty {
	var returns *DataTfEndpoint_KafkaSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) MessageFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) MessageMaxBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageMaxBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) NoHexPrefix() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"noHexPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) PartitionIncludeSchemaTable() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"partitionIncludeSchemaTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) SaslMechanism() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslMechanism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) SaslPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) SaslUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) SecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) SslCaCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) SslClientCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) SslClientKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) SslClientKeyPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) Topic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfEndpoint_KafkaSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataTfEndpoint_KafkaSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfEndpoint_KafkaSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.DataTfEndpoint.KafkaSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfEndpoint_KafkaSettingsPropertyOutputReference_Override(d DataTfEndpoint_KafkaSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.DataTfEndpoint.KafkaSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference)SetInternalValue(val *DataTfEndpoint_KafkaSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfEndpoint_KafkaSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

