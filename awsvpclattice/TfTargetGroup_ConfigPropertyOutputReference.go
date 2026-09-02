package awsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpclattice/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpclattice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTargetGroup_ConfigPropertyOutputReference interface {
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
	HealthCheck() TfTargetGroup_HealthCheckPropertyOutputReference
	// Experimental.
	HealthCheckInput() *TfTargetGroup_HealthCheckProperty
	// Experimental.
	InternalValue() *TfTargetGroup_ConfigProperty
	// Experimental.
	SetInternalValue(val *TfTargetGroup_ConfigProperty)
	// Experimental.
	IpAddressType() *string
	// Experimental.
	SetIpAddressType(val *string)
	// Experimental.
	IpAddressTypeInput() *string
	// Experimental.
	LambdaEventStructureVersion() *string
	// Experimental.
	SetLambdaEventStructureVersion(val *string)
	// Experimental.
	LambdaEventStructureVersionInput() *string
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
	ProtocolVersion() *string
	// Experimental.
	SetProtocolVersion(val *string)
	// Experimental.
	ProtocolVersionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VpcIdentifier() *string
	// Experimental.
	SetVpcIdentifier(val *string)
	// Experimental.
	VpcIdentifierInput() *string
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
	PutHealthCheck(value *TfTargetGroup_HealthCheckProperty)
	// Experimental.
	ResetHealthCheck()
	// Experimental.
	ResetIpAddressType()
	// Experimental.
	ResetLambdaEventStructureVersion()
	// Experimental.
	ResetPort()
	// Experimental.
	ResetProtocol()
	// Experimental.
	ResetProtocolVersion()
	// Experimental.
	ResetVpcIdentifier()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTargetGroup_ConfigPropertyOutputReference
type jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) HealthCheck() TfTargetGroup_HealthCheckPropertyOutputReference {
	var returns TfTargetGroup_HealthCheckPropertyOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) HealthCheckInput() *TfTargetGroup_HealthCheckProperty {
	var returns *TfTargetGroup_HealthCheckProperty
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) InternalValue() *TfTargetGroup_ConfigProperty {
	var returns *TfTargetGroup_ConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) LambdaEventStructureVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaEventStructureVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) LambdaEventStructureVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaEventStructureVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) ProtocolVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) ProtocolVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) VpcIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) VpcIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdentifierInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTargetGroup_ConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTargetGroup_ConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTargetGroup_ConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc-lattice.TfTargetGroup.ConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTargetGroup_ConfigPropertyOutputReference_Override(t TfTargetGroup_ConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc-lattice.TfTargetGroup.ConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference)SetInternalValue(val *TfTargetGroup_ConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference)SetLambdaEventStructureVersion(val *string) {
	if err := j.validateSetLambdaEventStructureVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaEventStructureVersion",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference)SetProtocolVersion(val *string) {
	if err := j.validateSetProtocolVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolVersion",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference)SetVpcIdentifier(val *string) {
	if err := j.validateSetVpcIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcIdentifier",
		val,
	)
}

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) PutHealthCheck(value *TfTargetGroup_HealthCheckProperty) {
	if err := t.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		t,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		t,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) ResetLambdaEventStructureVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaEventStructureVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		t,
		"resetPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		t,
		"resetProtocol",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) ResetProtocolVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetProtocolVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) ResetVpcIdentifier() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcIdentifier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTargetGroup_ConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

