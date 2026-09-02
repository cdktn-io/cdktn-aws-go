package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserPool_PasswordPolicyPropertyOutputReference interface {
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
	InternalValue() *TfUserPool_PasswordPolicyProperty
	// Experimental.
	SetInternalValue(val *TfUserPool_PasswordPolicyProperty)
	// Experimental.
	MinimumLength() *float64
	// Experimental.
	SetMinimumLength(val *float64)
	// Experimental.
	MinimumLengthInput() *float64
	// Experimental.
	PasswordHistorySize() *float64
	// Experimental.
	SetPasswordHistorySize(val *float64)
	// Experimental.
	PasswordHistorySizeInput() *float64
	// Experimental.
	RequireLowercase() interface{}
	// Experimental.
	SetRequireLowercase(val interface{})
	// Experimental.
	RequireLowercaseInput() interface{}
	// Experimental.
	RequireNumbers() interface{}
	// Experimental.
	SetRequireNumbers(val interface{})
	// Experimental.
	RequireNumbersInput() interface{}
	// Experimental.
	RequireSymbols() interface{}
	// Experimental.
	SetRequireSymbols(val interface{})
	// Experimental.
	RequireSymbolsInput() interface{}
	// Experimental.
	RequireUppercase() interface{}
	// Experimental.
	SetRequireUppercase(val interface{})
	// Experimental.
	RequireUppercaseInput() interface{}
	// Experimental.
	TemporaryPasswordValidityDays() *float64
	// Experimental.
	SetTemporaryPasswordValidityDays(val *float64)
	// Experimental.
	TemporaryPasswordValidityDaysInput() *float64
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
	ResetMinimumLength()
	// Experimental.
	ResetPasswordHistorySize()
	// Experimental.
	ResetRequireLowercase()
	// Experimental.
	ResetRequireNumbers()
	// Experimental.
	ResetRequireSymbols()
	// Experimental.
	ResetRequireUppercase()
	// Experimental.
	ResetTemporaryPasswordValidityDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfUserPool_PasswordPolicyPropertyOutputReference
type jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) InternalValue() *TfUserPool_PasswordPolicyProperty {
	var returns *TfUserPool_PasswordPolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) MinimumLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) MinimumLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) PasswordHistorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordHistorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) PasswordHistorySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordHistorySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) RequireLowercase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLowercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) RequireLowercaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLowercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) RequireNumbers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) RequireNumbersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) RequireSymbols() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSymbols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) RequireSymbolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSymbolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) RequireUppercase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireUppercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) RequireUppercaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireUppercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) TemporaryPasswordValidityDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temporaryPasswordValidityDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) TemporaryPasswordValidityDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temporaryPasswordValidityDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfUserPool_PasswordPolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfUserPool_PasswordPolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfUserPool_PasswordPolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfUserPool.PasswordPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfUserPool_PasswordPolicyPropertyOutputReference_Override(t TfUserPool_PasswordPolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfUserPool.PasswordPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference)SetInternalValue(val *TfUserPool_PasswordPolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference)SetMinimumLength(val *float64) {
	if err := j.validateSetMinimumLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumLength",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference)SetPasswordHistorySize(val *float64) {
	if err := j.validateSetPasswordHistorySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordHistorySize",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference)SetRequireLowercase(val interface{}) {
	if err := j.validateSetRequireLowercaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireLowercase",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference)SetRequireNumbers(val interface{}) {
	if err := j.validateSetRequireNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireNumbers",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference)SetRequireSymbols(val interface{}) {
	if err := j.validateSetRequireSymbolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireSymbols",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference)SetRequireUppercase(val interface{}) {
	if err := j.validateSetRequireUppercaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireUppercase",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference)SetTemporaryPasswordValidityDays(val *float64) {
	if err := j.validateSetTemporaryPasswordValidityDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"temporaryPasswordValidityDays",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) ResetMinimumLength() {
	_jsii_.InvokeVoid(
		t,
		"resetMinimumLength",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) ResetPasswordHistorySize() {
	_jsii_.InvokeVoid(
		t,
		"resetPasswordHistorySize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) ResetRequireLowercase() {
	_jsii_.InvokeVoid(
		t,
		"resetRequireLowercase",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) ResetRequireNumbers() {
	_jsii_.InvokeVoid(
		t,
		"resetRequireNumbers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) ResetRequireSymbols() {
	_jsii_.InvokeVoid(
		t,
		"resetRequireSymbols",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) ResetRequireUppercase() {
	_jsii_.InvokeVoid(
		t,
		"resetRequireUppercase",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) ResetTemporaryPasswordValidityDays() {
	_jsii_.InvokeVoid(
		t,
		"resetTemporaryPasswordValidityDays",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfUserPool_PasswordPolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

