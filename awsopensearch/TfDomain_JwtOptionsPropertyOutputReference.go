package awsopensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsopensearch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsopensearch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_JwtOptionsPropertyOutputReference interface {
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
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDomain_JwtOptionsProperty
	// Experimental.
	SetInternalValue(val *TfDomain_JwtOptionsProperty)
	// Experimental.
	JwksUrl() *string
	// Experimental.
	SetJwksUrl(val *string)
	// Experimental.
	JwksUrlInput() *string
	// Experimental.
	PublicKey() *string
	// Experimental.
	SetPublicKey(val *string)
	// Experimental.
	PublicKeyInput() *string
	// Experimental.
	RolesKey() *string
	// Experimental.
	SetRolesKey(val *string)
	// Experimental.
	RolesKeyInput() *string
	// Experimental.
	SubjectKey() *string
	// Experimental.
	SetSubjectKey(val *string)
	// Experimental.
	SubjectKeyInput() *string
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
	ResetEnabled()
	// Experimental.
	ResetJwksUrl()
	// Experimental.
	ResetPublicKey()
	// Experimental.
	ResetRolesKey()
	// Experimental.
	ResetSubjectKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDomain_JwtOptionsPropertyOutputReference
type jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) InternalValue() *TfDomain_JwtOptionsProperty {
	var returns *TfDomain_JwtOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) JwksUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) JwksUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) PublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) RolesKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolesKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) RolesKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolesKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) SubjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) SubjectKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_JwtOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDomain_JwtOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_JwtOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-opensearch.TfDomain.JwtOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_JwtOptionsPropertyOutputReference_Override(t TfDomain_JwtOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-opensearch.TfDomain.JwtOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference)SetInternalValue(val *TfDomain_JwtOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference)SetJwksUrl(val *string) {
	if err := j.validateSetJwksUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwksUrl",
		val,
	)
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference)SetPublicKey(val *string) {
	if err := j.validateSetPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicKey",
		val,
	)
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference)SetRolesKey(val *string) {
	if err := j.validateSetRolesKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rolesKey",
		val,
	)
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference)SetSubjectKey(val *string) {
	if err := j.validateSetSubjectKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectKey",
		val,
	)
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) ResetJwksUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetJwksUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) ResetPublicKey() {
	_jsii_.InvokeVoid(
		t,
		"resetPublicKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) ResetRolesKey() {
	_jsii_.InvokeVoid(
		t,
		"resetRolesKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) ResetSubjectKey() {
	_jsii_.InvokeVoid(
		t,
		"resetSubjectKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDomain_JwtOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

