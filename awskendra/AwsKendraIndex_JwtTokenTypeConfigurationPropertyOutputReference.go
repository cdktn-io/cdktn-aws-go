package awskendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskendra/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskendra/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ClaimRegex() *string
	// Experimental.
	SetClaimRegex(val *string)
	// Experimental.
	ClaimRegexInput() *string
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
	GroupAttributeField() *string
	// Experimental.
	SetGroupAttributeField(val *string)
	// Experimental.
	GroupAttributeFieldInput() *string
	// Experimental.
	InternalValue() *AwsKendraIndex_JwtTokenTypeConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsKendraIndex_JwtTokenTypeConfigurationProperty)
	// Experimental.
	Issuer() *string
	// Experimental.
	SetIssuer(val *string)
	// Experimental.
	IssuerInput() *string
	// Experimental.
	KeyLocation() *string
	// Experimental.
	SetKeyLocation(val *string)
	// Experimental.
	KeyLocationInput() *string
	// Experimental.
	SecretsManagerArn() *string
	// Experimental.
	SetSecretsManagerArn(val *string)
	// Experimental.
	SecretsManagerArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Url() *string
	// Experimental.
	SetUrl(val *string)
	// Experimental.
	UrlInput() *string
	// Experimental.
	UserNameAttributeField() *string
	// Experimental.
	SetUserNameAttributeField(val *string)
	// Experimental.
	UserNameAttributeFieldInput() *string
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
	ResetClaimRegex()
	// Experimental.
	ResetGroupAttributeField()
	// Experimental.
	ResetIssuer()
	// Experimental.
	ResetSecretsManagerArn()
	// Experimental.
	ResetUrl()
	// Experimental.
	ResetUserNameAttributeField()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference
type jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) ClaimRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"claimRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) ClaimRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"claimRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) GroupAttributeField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupAttributeField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) GroupAttributeFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupAttributeFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) InternalValue() *AwsKendraIndex_JwtTokenTypeConfigurationProperty {
	var returns *AwsKendraIndex_JwtTokenTypeConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) KeyLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) KeyLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) SecretsManagerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) SecretsManagerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) UserNameAttributeField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameAttributeField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) UserNameAttributeFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameAttributeFieldInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kendra.AwsKendraIndex.JwtTokenTypeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference_Override(a AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kendra.AwsKendraIndex.JwtTokenTypeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference)SetClaimRegex(val *string) {
	if err := j.validateSetClaimRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"claimRegex",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference)SetGroupAttributeField(val *string) {
	if err := j.validateSetGroupAttributeFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupAttributeField",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference)SetInternalValue(val *AwsKendraIndex_JwtTokenTypeConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference)SetKeyLocation(val *string) {
	if err := j.validateSetKeyLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyLocation",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference)SetSecretsManagerArn(val *string) {
	if err := j.validateSetSecretsManagerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerArn",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference)SetUserNameAttributeField(val *string) {
	if err := j.validateSetUserNameAttributeFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userNameAttributeField",
		val,
	)
}

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) ResetClaimRegex() {
	_jsii_.InvokeVoid(
		a,
		"resetClaimRegex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) ResetGroupAttributeField() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupAttributeField",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		a,
		"resetIssuer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) ResetSecretsManagerArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretsManagerArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) ResetUserNameAttributeField() {
	_jsii_.InvokeVoid(
		a,
		"resetUserNameAttributeField",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsKendraIndex_JwtTokenTypeConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

