package awsssoadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsssoadmin/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsssoadmin/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ClaimAttributePath() *string
	// Experimental.
	SetClaimAttributePath(val *string)
	// Experimental.
	ClaimAttributePathInput() *string
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
	IdentityStoreAttributePath() *string
	// Experimental.
	SetIdentityStoreAttributePath(val *string)
	// Experimental.
	IdentityStoreAttributePathInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IssuerUrl() *string
	// Experimental.
	SetIssuerUrl(val *string)
	// Experimental.
	IssuerUrlInput() *string
	// Experimental.
	JwksRetrievalOption() *string
	// Experimental.
	SetJwksRetrievalOption(val *string)
	// Experimental.
	JwksRetrievalOptionInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference
type jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) ClaimAttributePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"claimAttributePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) ClaimAttributePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"claimAttributePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) IdentityStoreAttributePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityStoreAttributePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) IdentityStoreAttributePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityStoreAttributePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) IssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) IssuerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) JwksRetrievalOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksRetrievalOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) JwksRetrievalOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksRetrievalOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sso-admin.TfTrustedTokenIssuer.OidcJwtConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference_Override(t TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sso-admin.TfTrustedTokenIssuer.OidcJwtConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference)SetClaimAttributePath(val *string) {
	if err := j.validateSetClaimAttributePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"claimAttributePath",
		val,
	)
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference)SetIdentityStoreAttributePath(val *string) {
	if err := j.validateSetIdentityStoreAttributePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityStoreAttributePath",
		val,
	)
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference)SetIssuerUrl(val *string) {
	if err := j.validateSetIssuerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerUrl",
		val,
	)
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference)SetJwksRetrievalOption(val *string) {
	if err := j.validateSetJwksRetrievalOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwksRetrievalOption",
		val,
	)
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTrustedTokenIssuer_OidcJwtConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

