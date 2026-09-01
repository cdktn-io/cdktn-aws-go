package awsfinspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfinspace/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfinspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
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
	IcmpTypeCode() AwsFinspaceKxEnvironment_IcmpTypeCodePropertyOutputReference
	// Experimental.
	IcmpTypeCodeInput() *AwsFinspaceKxEnvironment_IcmpTypeCodeProperty
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PortRange() AwsFinspaceKxEnvironment_PortRangePropertyOutputReference
	// Experimental.
	PortRangeInput() *AwsFinspaceKxEnvironment_PortRangeProperty
	// Experimental.
	Protocol() *string
	// Experimental.
	SetProtocol(val *string)
	// Experimental.
	ProtocolInput() *string
	// Experimental.
	RuleAction() *string
	// Experimental.
	SetRuleAction(val *string)
	// Experimental.
	RuleActionInput() *string
	// Experimental.
	RuleNumber() *float64
	// Experimental.
	SetRuleNumber(val *float64)
	// Experimental.
	RuleNumberInput() *float64
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
	PutIcmpTypeCode(value *AwsFinspaceKxEnvironment_IcmpTypeCodeProperty)
	// Experimental.
	PutPortRange(value *AwsFinspaceKxEnvironment_PortRangeProperty)
	// Experimental.
	ResetIcmpTypeCode()
	// Experimental.
	ResetPortRange()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference
type jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) IcmpTypeCode() AwsFinspaceKxEnvironment_IcmpTypeCodePropertyOutputReference {
	var returns AwsFinspaceKxEnvironment_IcmpTypeCodePropertyOutputReference
	_jsii_.Get(
		j,
		"icmpTypeCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) IcmpTypeCodeInput() *AwsFinspaceKxEnvironment_IcmpTypeCodeProperty {
	var returns *AwsFinspaceKxEnvironment_IcmpTypeCodeProperty
	_jsii_.Get(
		j,
		"icmpTypeCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) PortRange() AwsFinspaceKxEnvironment_PortRangePropertyOutputReference {
	var returns AwsFinspaceKxEnvironment_PortRangePropertyOutputReference
	_jsii_.Get(
		j,
		"portRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) PortRangeInput() *AwsFinspaceKxEnvironment_PortRangeProperty {
	var returns *AwsFinspaceKxEnvironment_PortRangeProperty
	_jsii_.Get(
		j,
		"portRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) RuleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) RuleActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) RuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) RuleNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-finspace.AwsFinspaceKxEnvironment.AttachmentNetworkAclConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference_Override(a AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-finspace.AwsFinspaceKxEnvironment.AttachmentNetworkAclConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference)SetCidrBlock(val *string) {
	if err := j.validateSetCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrBlock",
		val,
	)
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference)SetRuleAction(val *string) {
	if err := j.validateSetRuleActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleAction",
		val,
	)
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference)SetRuleNumber(val *float64) {
	if err := j.validateSetRuleNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleNumber",
		val,
	)
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) PutIcmpTypeCode(value *AwsFinspaceKxEnvironment_IcmpTypeCodeProperty) {
	if err := a.validatePutIcmpTypeCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIcmpTypeCode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) PutPortRange(value *AwsFinspaceKxEnvironment_PortRangeProperty) {
	if err := a.validatePutPortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPortRange",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) ResetIcmpTypeCode() {
	_jsii_.InvokeVoid(
		a,
		"resetIcmpTypeCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) ResetPortRange() {
	_jsii_.InvokeVoid(
		a,
		"resetPortRange",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

