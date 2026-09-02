package awsverifiedaccess

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsverifiedaccess/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsverifiedaccess/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpoint_RdsOptionsPropertyOutputReference interface {
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
	InternalValue() *TfEndpoint_RdsOptionsProperty
	// Experimental.
	SetInternalValue(val *TfEndpoint_RdsOptionsProperty)
	// Experimental.
	Port() *float64
	// Experimental.
	SetPort(val *float64)
	// Experimental.
	PortInput() *float64
	// Experimental.
	Protocol() *string
	// Experimental.
	SetProtocol(val *string)
	// Experimental.
	ProtocolInput() *string
	// Experimental.
	RdsDbClusterArn() *string
	// Experimental.
	SetRdsDbClusterArn(val *string)
	// Experimental.
	RdsDbClusterArnInput() *string
	// Experimental.
	RdsDbInstanceArn() *string
	// Experimental.
	SetRdsDbInstanceArn(val *string)
	// Experimental.
	RdsDbInstanceArnInput() *string
	// Experimental.
	RdsDbProxyArn() *string
	// Experimental.
	SetRdsDbProxyArn(val *string)
	// Experimental.
	RdsDbProxyArnInput() *string
	// Experimental.
	RdsEndpoint() *string
	// Experimental.
	SetRdsEndpoint(val *string)
	// Experimental.
	RdsEndpointInput() *string
	// Experimental.
	SubnetIds() *[]*string
	// Experimental.
	SetSubnetIds(val *[]*string)
	// Experimental.
	SubnetIdsInput() *[]*string
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
	ResetPort()
	// Experimental.
	ResetProtocol()
	// Experimental.
	ResetRdsDbClusterArn()
	// Experimental.
	ResetRdsDbInstanceArn()
	// Experimental.
	ResetRdsDbProxyArn()
	// Experimental.
	ResetRdsEndpoint()
	// Experimental.
	ResetSubnetIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEndpoint_RdsOptionsPropertyOutputReference
type jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) InternalValue() *TfEndpoint_RdsOptionsProperty {
	var returns *TfEndpoint_RdsOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) RdsDbClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) RdsDbClusterArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbClusterArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) RdsDbInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) RdsDbInstanceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbInstanceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) RdsDbProxyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbProxyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) RdsDbProxyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbProxyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) RdsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) RdsEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEndpoint_RdsOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEndpoint_RdsOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEndpoint_RdsOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-verified-access.TfEndpoint.RdsOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEndpoint_RdsOptionsPropertyOutputReference_Override(t TfEndpoint_RdsOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-verified-access.TfEndpoint.RdsOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference)SetInternalValue(val *TfEndpoint_RdsOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference)SetRdsDbClusterArn(val *string) {
	if err := j.validateSetRdsDbClusterArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsDbClusterArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference)SetRdsDbInstanceArn(val *string) {
	if err := j.validateSetRdsDbInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsDbInstanceArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference)SetRdsDbProxyArn(val *string) {
	if err := j.validateSetRdsDbProxyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsDbProxyArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference)SetRdsEndpoint(val *string) {
	if err := j.validateSetRdsEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsEndpoint",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		t,
		"resetPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		t,
		"resetProtocol",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) ResetRdsDbClusterArn() {
	_jsii_.InvokeVoid(
		t,
		"resetRdsDbClusterArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) ResetRdsDbInstanceArn() {
	_jsii_.InvokeVoid(
		t,
		"resetRdsDbInstanceArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) ResetRdsDbProxyArn() {
	_jsii_.InvokeVoid(
		t,
		"resetRdsDbProxyArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) ResetRdsEndpoint() {
	_jsii_.InvokeVoid(
		t,
		"resetRdsEndpoint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		t,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEndpoint_RdsOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

