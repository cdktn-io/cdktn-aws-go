package vpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/vpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/vpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDefaultNetworkAcl_EgressPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Action() *string
	// Experimental.
	SetAction(val *string)
	// Experimental.
	ActionInput() *string
	// Experimental.
	CidrBlock() *string
	// Experimental.
	SetCidrBlock(val *string)
	// Experimental.
	CidrBlockInput() *string
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
	FromPort() *float64
	// Experimental.
	SetFromPort(val *float64)
	// Experimental.
	FromPortInput() *float64
	// Experimental.
	IcmpCode() *float64
	// Experimental.
	SetIcmpCode(val *float64)
	// Experimental.
	IcmpCodeInput() *float64
	// Experimental.
	IcmpType() *float64
	// Experimental.
	SetIcmpType(val *float64)
	// Experimental.
	IcmpTypeInput() *float64
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Ipv6CidrBlock() *string
	// Experimental.
	SetIpv6CidrBlock(val *string)
	// Experimental.
	Ipv6CidrBlockInput() *string
	// Experimental.
	Protocol() *string
	// Experimental.
	SetProtocol(val *string)
	// Experimental.
	ProtocolInput() *string
	// Experimental.
	RuleNo() *float64
	// Experimental.
	SetRuleNo(val *float64)
	// Experimental.
	RuleNoInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ToPort() *float64
	// Experimental.
	SetToPort(val *float64)
	// Experimental.
	ToPortInput() *float64
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
	ResetCidrBlock()
	// Experimental.
	ResetIcmpCode()
	// Experimental.
	ResetIcmpType()
	// Experimental.
	ResetIpv6CidrBlock()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDefaultNetworkAcl_EgressPropertyOutputReference
type jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) FromPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) FromPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fromPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) IcmpCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"icmpCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) IcmpCodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"icmpCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) IcmpType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"icmpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) IcmpTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"icmpTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) Ipv6CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) Ipv6CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) RuleNo() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleNo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) RuleNoInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleNoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) ToPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) ToPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toPortInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDefaultNetworkAcl_EgressPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsDefaultNetworkAcl_EgressPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDefaultNetworkAcl_EgressPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsDefaultNetworkAcl.EgressPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDefaultNetworkAcl_EgressPropertyOutputReference_Override(a AwsDefaultNetworkAcl_EgressPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsDefaultNetworkAcl.EgressPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference)SetCidrBlock(val *string) {
	if err := j.validateSetCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrBlock",
		val,
	)
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference)SetFromPort(val *float64) {
	if err := j.validateSetFromPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fromPort",
		val,
	)
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference)SetIcmpCode(val *float64) {
	if err := j.validateSetIcmpCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icmpCode",
		val,
	)
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference)SetIcmpType(val *float64) {
	if err := j.validateSetIcmpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icmpType",
		val,
	)
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference)SetIpv6CidrBlock(val *string) {
	if err := j.validateSetIpv6CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6CidrBlock",
		val,
	)
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference)SetRuleNo(val *float64) {
	if err := j.validateSetRuleNoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleNo",
		val,
	)
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference)SetToPort(val *float64) {
	if err := j.validateSetToPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toPort",
		val,
	)
}

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) ResetCidrBlock() {
	_jsii_.InvokeVoid(
		a,
		"resetCidrBlock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) ResetIcmpCode() {
	_jsii_.InvokeVoid(
		a,
		"resetIcmpCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) ResetIcmpType() {
	_jsii_.InvokeVoid(
		a,
		"resetIcmpType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) ResetIpv6CidrBlock() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv6CidrBlock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDefaultNetworkAcl_EgressPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

