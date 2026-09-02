package awsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconfig/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConfigRule_CustomPolicyDetailsPropertyOutputReference interface {
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
	EnableDebugLogDelivery() interface{}
	// Experimental.
	SetEnableDebugLogDelivery(val interface{})
	// Experimental.
	EnableDebugLogDeliveryInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfConfigRule_CustomPolicyDetailsProperty
	// Experimental.
	SetInternalValue(val *TfConfigRule_CustomPolicyDetailsProperty)
	// Experimental.
	PolicyRuntime() *string
	// Experimental.
	SetPolicyRuntime(val *string)
	// Experimental.
	PolicyRuntimeInput() *string
	// Experimental.
	PolicyText() *string
	// Experimental.
	SetPolicyText(val *string)
	// Experimental.
	PolicyTextInput() *string
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
	ResetEnableDebugLogDelivery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfConfigRule_CustomPolicyDetailsPropertyOutputReference
type jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) EnableDebugLogDelivery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDebugLogDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) EnableDebugLogDeliveryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDebugLogDeliveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) InternalValue() *TfConfigRule_CustomPolicyDetailsProperty {
	var returns *TfConfigRule_CustomPolicyDetailsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) PolicyRuntime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) PolicyRuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) PolicyText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) PolicyTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConfigRule_CustomPolicyDetailsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConfigRule_CustomPolicyDetailsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConfigRule_CustomPolicyDetailsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-config.TfConfigRule.CustomPolicyDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConfigRule_CustomPolicyDetailsPropertyOutputReference_Override(t TfConfigRule_CustomPolicyDetailsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-config.TfConfigRule.CustomPolicyDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference)SetEnableDebugLogDelivery(val interface{}) {
	if err := j.validateSetEnableDebugLogDeliveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDebugLogDelivery",
		val,
	)
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference)SetInternalValue(val *TfConfigRule_CustomPolicyDetailsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference)SetPolicyRuntime(val *string) {
	if err := j.validateSetPolicyRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyRuntime",
		val,
	)
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference)SetPolicyText(val *string) {
	if err := j.validateSetPolicyTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyText",
		val,
	)
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) ResetEnableDebugLogDelivery() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableDebugLogDelivery",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfConfigRule_CustomPolicyDetailsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

