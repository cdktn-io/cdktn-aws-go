package awsvpnclient

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpnclient/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpnclient/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpoint_AuthenticationOptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ActiveDirectoryId() *string
	// Experimental.
	SetActiveDirectoryId(val *string)
	// Experimental.
	ActiveDirectoryIdInput() *string
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	RootCertificateChainArn() *string
	// Experimental.
	SetRootCertificateChainArn(val *string)
	// Experimental.
	RootCertificateChainArnInput() *string
	// Experimental.
	SamlProviderArn() *string
	// Experimental.
	SetSamlProviderArn(val *string)
	// Experimental.
	SamlProviderArnInput() *string
	// Experimental.
	SelfServiceSamlProviderArn() *string
	// Experimental.
	SetSelfServiceSamlProviderArn(val *string)
	// Experimental.
	SelfServiceSamlProviderArnInput() *string
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
	ResetActiveDirectoryId()
	// Experimental.
	ResetRootCertificateChainArn()
	// Experimental.
	ResetSamlProviderArn()
	// Experimental.
	ResetSelfServiceSamlProviderArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEndpoint_AuthenticationOptionsPropertyOutputReference
type jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) ActiveDirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) ActiveDirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) RootCertificateChainArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootCertificateChainArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) RootCertificateChainArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootCertificateChainArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) SamlProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) SamlProviderArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlProviderArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) SelfServiceSamlProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServiceSamlProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) SelfServiceSamlProviderArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServiceSamlProviderArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEndpoint_AuthenticationOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfEndpoint_AuthenticationOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEndpoint_AuthenticationOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpn-client.TfEndpoint.AuthenticationOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEndpoint_AuthenticationOptionsPropertyOutputReference_Override(t TfEndpoint_AuthenticationOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpn-client.TfEndpoint.AuthenticationOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference)SetActiveDirectoryId(val *string) {
	if err := j.validateSetActiveDirectoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDirectoryId",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference)SetRootCertificateChainArn(val *string) {
	if err := j.validateSetRootCertificateChainArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootCertificateChainArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference)SetSamlProviderArn(val *string) {
	if err := j.validateSetSamlProviderArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samlProviderArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference)SetSelfServiceSamlProviderArn(val *string) {
	if err := j.validateSetSelfServiceSamlProviderArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfServiceSamlProviderArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) ResetActiveDirectoryId() {
	_jsii_.InvokeVoid(
		t,
		"resetActiveDirectoryId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) ResetRootCertificateChainArn() {
	_jsii_.InvokeVoid(
		t,
		"resetRootCertificateChainArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) ResetSamlProviderArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSamlProviderArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) ResetSelfServiceSamlProviderArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSelfServiceSamlProviderArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEndpoint_AuthenticationOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

