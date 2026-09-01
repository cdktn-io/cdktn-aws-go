package awsfms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFmsPolicy_PolicyOptionPropertyOutputReference interface {
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
	InternalValue() *AwsFmsPolicy_PolicyOptionProperty
	// Experimental.
	SetInternalValue(val *AwsFmsPolicy_PolicyOptionProperty)
	// Experimental.
	NetworkAclCommonPolicy() AwsFmsPolicy_NetworkAclCommonPolicyPropertyOutputReference
	// Experimental.
	NetworkAclCommonPolicyInput() *AwsFmsPolicy_NetworkAclCommonPolicyProperty
	// Experimental.
	NetworkFirewallPolicy() AwsFmsPolicy_NetworkFirewallPolicyPropertyOutputReference
	// Experimental.
	NetworkFirewallPolicyInput() *AwsFmsPolicy_NetworkFirewallPolicyProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ThirdPartyFirewallPolicy() AwsFmsPolicy_ThirdPartyFirewallPolicyPropertyOutputReference
	// Experimental.
	ThirdPartyFirewallPolicyInput() *AwsFmsPolicy_ThirdPartyFirewallPolicyProperty
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
	PutNetworkAclCommonPolicy(value *AwsFmsPolicy_NetworkAclCommonPolicyProperty)
	// Experimental.
	PutNetworkFirewallPolicy(value *AwsFmsPolicy_NetworkFirewallPolicyProperty)
	// Experimental.
	PutThirdPartyFirewallPolicy(value *AwsFmsPolicy_ThirdPartyFirewallPolicyProperty)
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

// The jsii proxy struct for AwsFmsPolicy_PolicyOptionPropertyOutputReference
type jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) InternalValue() *AwsFmsPolicy_PolicyOptionProperty {
	var returns *AwsFmsPolicy_PolicyOptionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) NetworkAclCommonPolicy() AwsFmsPolicy_NetworkAclCommonPolicyPropertyOutputReference {
	var returns AwsFmsPolicy_NetworkAclCommonPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"networkAclCommonPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) NetworkAclCommonPolicyInput() *AwsFmsPolicy_NetworkAclCommonPolicyProperty {
	var returns *AwsFmsPolicy_NetworkAclCommonPolicyProperty
	_jsii_.Get(
		j,
		"networkAclCommonPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) NetworkFirewallPolicy() AwsFmsPolicy_NetworkFirewallPolicyPropertyOutputReference {
	var returns AwsFmsPolicy_NetworkFirewallPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"networkFirewallPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) NetworkFirewallPolicyInput() *AwsFmsPolicy_NetworkFirewallPolicyProperty {
	var returns *AwsFmsPolicy_NetworkFirewallPolicyProperty
	_jsii_.Get(
		j,
		"networkFirewallPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) ThirdPartyFirewallPolicy() AwsFmsPolicy_ThirdPartyFirewallPolicyPropertyOutputReference {
	var returns AwsFmsPolicy_ThirdPartyFirewallPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"thirdPartyFirewallPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) ThirdPartyFirewallPolicyInput() *AwsFmsPolicy_ThirdPartyFirewallPolicyProperty {
	var returns *AwsFmsPolicy_ThirdPartyFirewallPolicyProperty
	_jsii_.Get(
		j,
		"thirdPartyFirewallPolicyInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFmsPolicy_PolicyOptionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFmsPolicy_PolicyOptionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFmsPolicy_PolicyOptionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fms.AwsFmsPolicy.PolicyOptionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFmsPolicy_PolicyOptionPropertyOutputReference_Override(a AwsFmsPolicy_PolicyOptionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fms.AwsFmsPolicy.PolicyOptionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference)SetInternalValue(val *AwsFmsPolicy_PolicyOptionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) PutNetworkAclCommonPolicy(value *AwsFmsPolicy_NetworkAclCommonPolicyProperty) {
	if err := a.validatePutNetworkAclCommonPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkAclCommonPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) PutNetworkFirewallPolicy(value *AwsFmsPolicy_NetworkFirewallPolicyProperty) {
	if err := a.validatePutNetworkFirewallPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkFirewallPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) PutThirdPartyFirewallPolicy(value *AwsFmsPolicy_ThirdPartyFirewallPolicyProperty) {
	if err := a.validatePutThirdPartyFirewallPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putThirdPartyFirewallPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) ResetNetworkAclCommonPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkAclCommonPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) ResetNetworkFirewallPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkFirewallPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) ResetThirdPartyFirewallPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetThirdPartyFirewallPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFmsPolicy_PolicyOptionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

