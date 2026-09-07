package verifiedaccess

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/verifiedaccess/jsii"

	"github.com/cdktn-io/cdktn-aws-go/verifiedaccess/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEndpoint_RdsOptionsPropertyOutputReference interface {
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
	InternalValue() *AwsEndpoint_RdsOptionsProperty
	// Experimental.
	SetInternalValue(val *AwsEndpoint_RdsOptionsProperty)
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

// The jsii proxy struct for AwsEndpoint_RdsOptionsPropertyOutputReference
type jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) InternalValue() *AwsEndpoint_RdsOptionsProperty {
	var returns *AwsEndpoint_RdsOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) RdsDbClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) RdsDbClusterArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbClusterArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) RdsDbInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) RdsDbInstanceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbInstanceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) RdsDbProxyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbProxyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) RdsDbProxyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDbProxyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) RdsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) RdsEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEndpoint_RdsOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEndpoint_RdsOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEndpoint_RdsOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-verified-access.AwsEndpoint.RdsOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEndpoint_RdsOptionsPropertyOutputReference_Override(a AwsEndpoint_RdsOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-verified-access.AwsEndpoint.RdsOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference)SetInternalValue(val *AwsEndpoint_RdsOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference)SetRdsDbClusterArn(val *string) {
	if err := j.validateSetRdsDbClusterArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsDbClusterArn",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference)SetRdsDbInstanceArn(val *string) {
	if err := j.validateSetRdsDbInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsDbInstanceArn",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference)SetRdsDbProxyArn(val *string) {
	if err := j.validateSetRdsDbProxyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsDbProxyArn",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference)SetRdsEndpoint(val *string) {
	if err := j.validateSetRdsEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) ResetRdsDbClusterArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRdsDbClusterArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) ResetRdsDbInstanceArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRdsDbInstanceArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) ResetRdsDbProxyArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRdsDbProxyArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) ResetRdsEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetRdsEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEndpoint_RdsOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

