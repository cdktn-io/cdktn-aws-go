package msk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/msk/jsii"

	"github.com/cdktn-io/cdktn-aws-go/msk/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationProperty
	// Experimental.
	SetInternalValue(val *AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationProperty)
	// Experimental.
	Sasl() AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSaslPropertyOutputReference
	// Experimental.
	SaslInput() *AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSaslProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Tls() interface{}
	// Experimental.
	SetTls(val interface{})
	// Experimental.
	TlsInput() interface{}
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
	PutSasl(value *AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSaslProperty)
	// Experimental.
	ResetSasl()
	// Experimental.
	ResetTls()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference
type jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) InternalValue() *AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationProperty {
	var returns *AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) Sasl() AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSaslPropertyOutputReference {
	var returns AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSaslPropertyOutputReference
	_jsii_.Get(
		j,
		"sasl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) SaslInput() *AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSaslProperty {
	var returns *AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSaslProperty
	_jsii_.Get(
		j,
		"saslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) Tls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) TlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-msk.AwsCluster.BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference_Override(a AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-msk.AwsCluster.BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference)SetInternalValue(val *AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference)SetTls(val interface{}) {
	if err := j.validateSetTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tls",
		val,
	)
}

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) PutSasl(value *AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSaslProperty) {
	if err := a.validatePutSaslParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSasl",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) ResetSasl() {
	_jsii_.InvokeVoid(
		a,
		"resetSasl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		a,
		"resetTls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

