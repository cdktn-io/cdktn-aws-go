package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference interface {
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
	// Experimental.
	ContentSecurityPolicy() TfResponseHeadersPolicy_ContentSecurityPolicyPropertyOutputReference
	// Experimental.
	ContentSecurityPolicyInput() *TfResponseHeadersPolicy_ContentSecurityPolicyProperty
	// Experimental.
	ContentTypeOptions() TfResponseHeadersPolicy_ContentTypeOptionsPropertyOutputReference
	// Experimental.
	ContentTypeOptionsInput() *TfResponseHeadersPolicy_ContentTypeOptionsProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	FrameOptions() TfResponseHeadersPolicy_FrameOptionsPropertyOutputReference
	// Experimental.
	FrameOptionsInput() *TfResponseHeadersPolicy_FrameOptionsProperty
	// Experimental.
	InternalValue() *TfResponseHeadersPolicy_SecurityHeadersConfigProperty
	// Experimental.
	SetInternalValue(val *TfResponseHeadersPolicy_SecurityHeadersConfigProperty)
	// Experimental.
	ReferrerPolicy() TfResponseHeadersPolicy_ReferrerPolicyPropertyOutputReference
	// Experimental.
	ReferrerPolicyInput() *TfResponseHeadersPolicy_ReferrerPolicyProperty
	// Experimental.
	StrictTransportSecurity() TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference
	// Experimental.
	StrictTransportSecurityInput() *TfResponseHeadersPolicy_StrictTransportSecurityProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	XssProtection() TfResponseHeadersPolicy_XssProtectionPropertyOutputReference
	// Experimental.
	XssProtectionInput() *TfResponseHeadersPolicy_XssProtectionProperty
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
	PutContentSecurityPolicy(value *TfResponseHeadersPolicy_ContentSecurityPolicyProperty)
	// Experimental.
	PutContentTypeOptions(value *TfResponseHeadersPolicy_ContentTypeOptionsProperty)
	// Experimental.
	PutFrameOptions(value *TfResponseHeadersPolicy_FrameOptionsProperty)
	// Experimental.
	PutReferrerPolicy(value *TfResponseHeadersPolicy_ReferrerPolicyProperty)
	// Experimental.
	PutStrictTransportSecurity(value *TfResponseHeadersPolicy_StrictTransportSecurityProperty)
	// Experimental.
	PutXssProtection(value *TfResponseHeadersPolicy_XssProtectionProperty)
	// Experimental.
	ResetContentSecurityPolicy()
	// Experimental.
	ResetContentTypeOptions()
	// Experimental.
	ResetFrameOptions()
	// Experimental.
	ResetReferrerPolicy()
	// Experimental.
	ResetStrictTransportSecurity()
	// Experimental.
	ResetXssProtection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference
type jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ContentSecurityPolicy() TfResponseHeadersPolicy_ContentSecurityPolicyPropertyOutputReference {
	var returns TfResponseHeadersPolicy_ContentSecurityPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"contentSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ContentSecurityPolicyInput() *TfResponseHeadersPolicy_ContentSecurityPolicyProperty {
	var returns *TfResponseHeadersPolicy_ContentSecurityPolicyProperty
	_jsii_.Get(
		j,
		"contentSecurityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ContentTypeOptions() TfResponseHeadersPolicy_ContentTypeOptionsPropertyOutputReference {
	var returns TfResponseHeadersPolicy_ContentTypeOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"contentTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ContentTypeOptionsInput() *TfResponseHeadersPolicy_ContentTypeOptionsProperty {
	var returns *TfResponseHeadersPolicy_ContentTypeOptionsProperty
	_jsii_.Get(
		j,
		"contentTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) FrameOptions() TfResponseHeadersPolicy_FrameOptionsPropertyOutputReference {
	var returns TfResponseHeadersPolicy_FrameOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"frameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) FrameOptionsInput() *TfResponseHeadersPolicy_FrameOptionsProperty {
	var returns *TfResponseHeadersPolicy_FrameOptionsProperty
	_jsii_.Get(
		j,
		"frameOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) InternalValue() *TfResponseHeadersPolicy_SecurityHeadersConfigProperty {
	var returns *TfResponseHeadersPolicy_SecurityHeadersConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ReferrerPolicy() TfResponseHeadersPolicy_ReferrerPolicyPropertyOutputReference {
	var returns TfResponseHeadersPolicy_ReferrerPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"referrerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ReferrerPolicyInput() *TfResponseHeadersPolicy_ReferrerPolicyProperty {
	var returns *TfResponseHeadersPolicy_ReferrerPolicyProperty
	_jsii_.Get(
		j,
		"referrerPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) StrictTransportSecurity() TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference {
	var returns TfResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference
	_jsii_.Get(
		j,
		"strictTransportSecurity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) StrictTransportSecurityInput() *TfResponseHeadersPolicy_StrictTransportSecurityProperty {
	var returns *TfResponseHeadersPolicy_StrictTransportSecurityProperty
	_jsii_.Get(
		j,
		"strictTransportSecurityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) XssProtection() TfResponseHeadersPolicy_XssProtectionPropertyOutputReference {
	var returns TfResponseHeadersPolicy_XssProtectionPropertyOutputReference
	_jsii_.Get(
		j,
		"xssProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) XssProtectionInput() *TfResponseHeadersPolicy_XssProtectionProperty {
	var returns *TfResponseHeadersPolicy_XssProtectionProperty
	_jsii_.Get(
		j,
		"xssProtectionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfResponseHeadersPolicy.SecurityHeadersConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference_Override(t TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfResponseHeadersPolicy.SecurityHeadersConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference)SetInternalValue(val *TfResponseHeadersPolicy_SecurityHeadersConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutContentSecurityPolicy(value *TfResponseHeadersPolicy_ContentSecurityPolicyProperty) {
	if err := t.validatePutContentSecurityPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putContentSecurityPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutContentTypeOptions(value *TfResponseHeadersPolicy_ContentTypeOptionsProperty) {
	if err := t.validatePutContentTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putContentTypeOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutFrameOptions(value *TfResponseHeadersPolicy_FrameOptionsProperty) {
	if err := t.validatePutFrameOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFrameOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutReferrerPolicy(value *TfResponseHeadersPolicy_ReferrerPolicyProperty) {
	if err := t.validatePutReferrerPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putReferrerPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutStrictTransportSecurity(value *TfResponseHeadersPolicy_StrictTransportSecurityProperty) {
	if err := t.validatePutStrictTransportSecurityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStrictTransportSecurity",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutXssProtection(value *TfResponseHeadersPolicy_XssProtectionProperty) {
	if err := t.validatePutXssProtectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putXssProtection",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetContentSecurityPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetContentSecurityPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetContentTypeOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetContentTypeOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetFrameOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetFrameOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetReferrerPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetReferrerPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetStrictTransportSecurity() {
	_jsii_.InvokeVoid(
		t,
		"resetStrictTransportSecurity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetXssProtection() {
	_jsii_.InvokeVoid(
		t,
		"resetXssProtection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

