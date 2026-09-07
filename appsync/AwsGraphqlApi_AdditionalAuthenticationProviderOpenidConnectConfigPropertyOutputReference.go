package appsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appsync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appsync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthTtl() *float64
	// Experimental.
	SetAuthTtl(val *float64)
	// Experimental.
	AuthTtlInput() *float64
	// Experimental.
	ClientId() *string
	// Experimental.
	SetClientId(val *string)
	// Experimental.
	ClientIdInput() *string
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
	IatTtl() *float64
	// Experimental.
	SetIatTtl(val *float64)
	// Experimental.
	IatTtlInput() *float64
	// Experimental.
	InternalValue() *AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty
	// Experimental.
	SetInternalValue(val *AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty)
	// Experimental.
	Issuer() *string
	// Experimental.
	SetIssuer(val *string)
	// Experimental.
	IssuerInput() *string
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
	ResetAuthTtl()
	// Experimental.
	ResetClientId()
	// Experimental.
	ResetIatTtl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference
type jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) AuthTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) AuthTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) IatTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iatTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) IatTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iatTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) InternalValue() *AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty {
	var returns *AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsGraphqlApi.AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference_Override(a AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsGraphqlApi.AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference)SetAuthTtl(val *float64) {
	if err := j.validateSetAuthTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authTtl",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference)SetIatTtl(val *float64) {
	if err := j.validateSetIatTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iatTtl",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference)SetInternalValue(val *AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) ResetAuthTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) ResetIatTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetIatTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

