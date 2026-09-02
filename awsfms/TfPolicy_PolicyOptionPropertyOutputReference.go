package awsfms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPolicy_PolicyOptionPropertyOutputReference interface {
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
	InternalValue() *TfPolicy_PolicyOptionProperty
	// Experimental.
	SetInternalValue(val *TfPolicy_PolicyOptionProperty)
	// Experimental.
	NetworkAclCommonPolicy() TfPolicy_NetworkAclCommonPolicyPropertyOutputReference
	// Experimental.
	NetworkAclCommonPolicyInput() *TfPolicy_NetworkAclCommonPolicyProperty
	// Experimental.
	NetworkFirewallPolicy() TfPolicy_NetworkFirewallPolicyPropertyOutputReference
	// Experimental.
	NetworkFirewallPolicyInput() *TfPolicy_NetworkFirewallPolicyProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ThirdPartyFirewallPolicy() TfPolicy_ThirdPartyFirewallPolicyPropertyOutputReference
	// Experimental.
	ThirdPartyFirewallPolicyInput() *TfPolicy_ThirdPartyFirewallPolicyProperty
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
	PutNetworkAclCommonPolicy(value *TfPolicy_NetworkAclCommonPolicyProperty)
	// Experimental.
	PutNetworkFirewallPolicy(value *TfPolicy_NetworkFirewallPolicyProperty)
	// Experimental.
	PutThirdPartyFirewallPolicy(value *TfPolicy_ThirdPartyFirewallPolicyProperty)
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

// The jsii proxy struct for TfPolicy_PolicyOptionPropertyOutputReference
type jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) InternalValue() *TfPolicy_PolicyOptionProperty {
	var returns *TfPolicy_PolicyOptionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) NetworkAclCommonPolicy() TfPolicy_NetworkAclCommonPolicyPropertyOutputReference {
	var returns TfPolicy_NetworkAclCommonPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"networkAclCommonPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) NetworkAclCommonPolicyInput() *TfPolicy_NetworkAclCommonPolicyProperty {
	var returns *TfPolicy_NetworkAclCommonPolicyProperty
	_jsii_.Get(
		j,
		"networkAclCommonPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) NetworkFirewallPolicy() TfPolicy_NetworkFirewallPolicyPropertyOutputReference {
	var returns TfPolicy_NetworkFirewallPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"networkFirewallPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) NetworkFirewallPolicyInput() *TfPolicy_NetworkFirewallPolicyProperty {
	var returns *TfPolicy_NetworkFirewallPolicyProperty
	_jsii_.Get(
		j,
		"networkFirewallPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) ThirdPartyFirewallPolicy() TfPolicy_ThirdPartyFirewallPolicyPropertyOutputReference {
	var returns TfPolicy_ThirdPartyFirewallPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"thirdPartyFirewallPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) ThirdPartyFirewallPolicyInput() *TfPolicy_ThirdPartyFirewallPolicyProperty {
	var returns *TfPolicy_ThirdPartyFirewallPolicyProperty
	_jsii_.Get(
		j,
		"thirdPartyFirewallPolicyInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPolicy_PolicyOptionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPolicy_PolicyOptionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPolicy_PolicyOptionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fms.TfPolicy.PolicyOptionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPolicy_PolicyOptionPropertyOutputReference_Override(t TfPolicy_PolicyOptionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fms.TfPolicy.PolicyOptionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference)SetInternalValue(val *TfPolicy_PolicyOptionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) PutNetworkAclCommonPolicy(value *TfPolicy_NetworkAclCommonPolicyProperty) {
	if err := t.validatePutNetworkAclCommonPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkAclCommonPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) PutNetworkFirewallPolicy(value *TfPolicy_NetworkFirewallPolicyProperty) {
	if err := t.validatePutNetworkFirewallPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkFirewallPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) PutThirdPartyFirewallPolicy(value *TfPolicy_ThirdPartyFirewallPolicyProperty) {
	if err := t.validatePutThirdPartyFirewallPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putThirdPartyFirewallPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) ResetNetworkAclCommonPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkAclCommonPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) ResetNetworkFirewallPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkFirewallPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) ResetThirdPartyFirewallPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetThirdPartyFirewallPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPolicy_PolicyOptionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

