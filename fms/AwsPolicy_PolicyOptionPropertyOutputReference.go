package fms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/fms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/fms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPolicy_PolicyOptionPropertyOutputReference interface {
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
	InternalValue() *AwsPolicy_PolicyOptionProperty
	// Experimental.
	SetInternalValue(val *AwsPolicy_PolicyOptionProperty)
	// Experimental.
	NetworkAclCommonPolicy() AwsPolicy_NetworkAclCommonPolicyPropertyOutputReference
	// Experimental.
	NetworkAclCommonPolicyInput() *AwsPolicy_NetworkAclCommonPolicyProperty
	// Experimental.
	NetworkFirewallPolicy() AwsPolicy_NetworkFirewallPolicyPropertyOutputReference
	// Experimental.
	NetworkFirewallPolicyInput() *AwsPolicy_NetworkFirewallPolicyProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ThirdPartyFirewallPolicy() AwsPolicy_ThirdPartyFirewallPolicyPropertyOutputReference
	// Experimental.
	ThirdPartyFirewallPolicyInput() *AwsPolicy_ThirdPartyFirewallPolicyProperty
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
	PutNetworkAclCommonPolicy(value *AwsPolicy_NetworkAclCommonPolicyProperty)
	// Experimental.
	PutNetworkFirewallPolicy(value *AwsPolicy_NetworkFirewallPolicyProperty)
	// Experimental.
	PutThirdPartyFirewallPolicy(value *AwsPolicy_ThirdPartyFirewallPolicyProperty)
	// Experimental.
	ResetNetworkAclCommonPolicy()
	// Experimental.
	ResetNetworkFirewallPolicy()
	// Experimental.
	ResetThirdPartyFirewallPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsPolicy_PolicyOptionPropertyOutputReference
type jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) InternalValue() *AwsPolicy_PolicyOptionProperty {
	var returns *AwsPolicy_PolicyOptionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) NetworkAclCommonPolicy() AwsPolicy_NetworkAclCommonPolicyPropertyOutputReference {
	var returns AwsPolicy_NetworkAclCommonPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"networkAclCommonPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) NetworkAclCommonPolicyInput() *AwsPolicy_NetworkAclCommonPolicyProperty {
	var returns *AwsPolicy_NetworkAclCommonPolicyProperty
	_jsii_.Get(
		j,
		"networkAclCommonPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) NetworkFirewallPolicy() AwsPolicy_NetworkFirewallPolicyPropertyOutputReference {
	var returns AwsPolicy_NetworkFirewallPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"networkFirewallPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) NetworkFirewallPolicyInput() *AwsPolicy_NetworkFirewallPolicyProperty {
	var returns *AwsPolicy_NetworkFirewallPolicyProperty
	_jsii_.Get(
		j,
		"networkFirewallPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) ThirdPartyFirewallPolicy() AwsPolicy_ThirdPartyFirewallPolicyPropertyOutputReference {
	var returns AwsPolicy_ThirdPartyFirewallPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"thirdPartyFirewallPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) ThirdPartyFirewallPolicyInput() *AwsPolicy_ThirdPartyFirewallPolicyProperty {
	var returns *AwsPolicy_ThirdPartyFirewallPolicyProperty
	_jsii_.Get(
		j,
		"thirdPartyFirewallPolicyInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPolicy_PolicyOptionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPolicy_PolicyOptionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPolicy_PolicyOptionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fms.AwsPolicy.PolicyOptionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPolicy_PolicyOptionPropertyOutputReference_Override(a AwsPolicy_PolicyOptionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fms.AwsPolicy.PolicyOptionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference)SetInternalValue(val *AwsPolicy_PolicyOptionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) PutNetworkAclCommonPolicy(value *AwsPolicy_NetworkAclCommonPolicyProperty) {
	if err := a.validatePutNetworkAclCommonPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkAclCommonPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) PutNetworkFirewallPolicy(value *AwsPolicy_NetworkFirewallPolicyProperty) {
	if err := a.validatePutNetworkFirewallPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkFirewallPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) PutThirdPartyFirewallPolicy(value *AwsPolicy_ThirdPartyFirewallPolicyProperty) {
	if err := a.validatePutThirdPartyFirewallPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putThirdPartyFirewallPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) ResetNetworkAclCommonPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkAclCommonPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) ResetNetworkFirewallPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkFirewallPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) ResetThirdPartyFirewallPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetThirdPartyFirewallPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPolicy_PolicyOptionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

