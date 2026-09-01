package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCognitoUserPool_PasswordPolicyPropertyOutputReference interface {
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
	InternalValue() *AwsCognitoUserPool_PasswordPolicyProperty
	// Experimental.
	SetInternalValue(val *AwsCognitoUserPool_PasswordPolicyProperty)
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

// The jsii proxy struct for AwsCognitoUserPool_PasswordPolicyPropertyOutputReference
type jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) InternalValue() *AwsCognitoUserPool_PasswordPolicyProperty {
	var returns *AwsCognitoUserPool_PasswordPolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) MinimumLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) MinimumLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) PasswordHistorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordHistorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) PasswordHistorySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordHistorySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) RequireLowercase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLowercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) RequireLowercaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLowercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) RequireNumbers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) RequireNumbersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) RequireSymbols() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSymbols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) RequireSymbolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSymbolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) RequireUppercase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireUppercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) RequireUppercaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireUppercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) TemporaryPasswordValidityDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temporaryPasswordValidityDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) TemporaryPasswordValidityDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temporaryPasswordValidityDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCognitoUserPool_PasswordPolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCognitoUserPool_PasswordPolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCognitoUserPool_PasswordPolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoUserPool.PasswordPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCognitoUserPool_PasswordPolicyPropertyOutputReference_Override(a AwsCognitoUserPool_PasswordPolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoUserPool.PasswordPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference)SetInternalValue(val *AwsCognitoUserPool_PasswordPolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference)SetMinimumLength(val *float64) {
	if err := j.validateSetMinimumLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumLength",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference)SetPasswordHistorySize(val *float64) {
	if err := j.validateSetPasswordHistorySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordHistorySize",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference)SetRequireLowercase(val interface{}) {
	if err := j.validateSetRequireLowercaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireLowercase",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference)SetRequireNumbers(val interface{}) {
	if err := j.validateSetRequireNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireNumbers",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference)SetRequireSymbols(val interface{}) {
	if err := j.validateSetRequireSymbolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireSymbols",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference)SetRequireUppercase(val interface{}) {
	if err := j.validateSetRequireUppercaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireUppercase",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference)SetTemporaryPasswordValidityDays(val *float64) {
	if err := j.validateSetTemporaryPasswordValidityDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"temporaryPasswordValidityDays",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) ResetMinimumLength() {
	_jsii_.InvokeVoid(
		a,
		"resetMinimumLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) ResetPasswordHistorySize() {
	_jsii_.InvokeVoid(
		a,
		"resetPasswordHistorySize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) ResetRequireLowercase() {
	_jsii_.InvokeVoid(
		a,
		"resetRequireLowercase",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) ResetRequireNumbers() {
	_jsii_.InvokeVoid(
		a,
		"resetRequireNumbers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) ResetRequireSymbols() {
	_jsii_.InvokeVoid(
		a,
		"resetRequireSymbols",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) ResetRequireUppercase() {
	_jsii_.InvokeVoid(
		a,
		"resetRequireUppercase",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) ResetTemporaryPasswordValidityDays() {
	_jsii_.InvokeVoid(
		a,
		"resetTemporaryPasswordValidityDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCognitoUserPool_PasswordPolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

