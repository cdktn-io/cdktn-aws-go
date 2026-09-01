package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrock/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrock/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference interface {
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
	InputAction() *string
	// Experimental.
	SetInputAction(val *string)
	// Experimental.
	InputActionInput() *string
	// Experimental.
	InputEnabled() interface{}
	// Experimental.
	SetInputEnabled(val interface{})
	// Experimental.
	InputEnabledInput() interface{}
	// Experimental.
	InputModalities() *[]*string
	// Experimental.
	SetInputModalities(val *[]*string)
	// Experimental.
	InputModalitiesInput() *[]*string
	// Experimental.
	InputStrength() *string
	// Experimental.
	SetInputStrength(val *string)
	// Experimental.
	InputStrengthInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	OutputAction() *string
	// Experimental.
	SetOutputAction(val *string)
	// Experimental.
	OutputActionInput() *string
	// Experimental.
	OutputEnabled() interface{}
	// Experimental.
	SetOutputEnabled(val interface{})
	// Experimental.
	OutputEnabledInput() interface{}
	// Experimental.
	OutputModalities() *[]*string
	// Experimental.
	SetOutputModalities(val *[]*string)
	// Experimental.
	OutputModalitiesInput() *[]*string
	// Experimental.
	OutputStrength() *string
	// Experimental.
	SetOutputStrength(val *string)
	// Experimental.
	OutputStrengthInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	ResetInputAction()
	// Experimental.
	ResetInputEnabled()
	// Experimental.
	ResetInputModalities()
	// Experimental.
	ResetOutputAction()
	// Experimental.
	ResetOutputEnabled()
	// Experimental.
	ResetOutputModalities()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference
type jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InputAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InputActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InputEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InputEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InputModalities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputModalities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InputModalitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputModalitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InputStrength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputStrength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InputStrengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputStrengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) OutputAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) OutputActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) OutputEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) OutputEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) OutputModalities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputModalities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) OutputModalitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputModalitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) OutputStrength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputStrength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) OutputStrengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputStrengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock.AwsBedrockGuardrail.ContentPolicyConfigFiltersConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference_Override(a AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock.AwsBedrockGuardrail.ContentPolicyConfigFiltersConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetInputAction(val *string) {
	if err := j.validateSetInputActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputAction",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetInputEnabled(val interface{}) {
	if err := j.validateSetInputEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetInputModalities(val *[]*string) {
	if err := j.validateSetInputModalitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputModalities",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetInputStrength(val *string) {
	if err := j.validateSetInputStrengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputStrength",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetOutputAction(val *string) {
	if err := j.validateSetOutputActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputAction",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetOutputEnabled(val interface{}) {
	if err := j.validateSetOutputEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetOutputModalities(val *[]*string) {
	if err := j.validateSetOutputModalitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputModalities",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetOutputStrength(val *string) {
	if err := j.validateSetOutputStrengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputStrength",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ResetInputAction() {
	_jsii_.InvokeVoid(
		a,
		"resetInputAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ResetInputEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetInputEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ResetInputModalities() {
	_jsii_.InvokeVoid(
		a,
		"resetInputModalities",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ResetOutputAction() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ResetOutputEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ResetOutputModalities() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputModalities",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBedrockGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

