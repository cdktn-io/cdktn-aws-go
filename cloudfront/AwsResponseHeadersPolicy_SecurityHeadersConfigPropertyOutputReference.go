package cloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference interface {
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
	ContentSecurityPolicy() AwsResponseHeadersPolicy_ContentSecurityPolicyPropertyOutputReference
	// Experimental.
	ContentSecurityPolicyInput() *AwsResponseHeadersPolicy_ContentSecurityPolicyProperty
	// Experimental.
	ContentTypeOptions() AwsResponseHeadersPolicy_ContentTypeOptionsPropertyOutputReference
	// Experimental.
	ContentTypeOptionsInput() *AwsResponseHeadersPolicy_ContentTypeOptionsProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	FrameOptions() AwsResponseHeadersPolicy_FrameOptionsPropertyOutputReference
	// Experimental.
	FrameOptionsInput() *AwsResponseHeadersPolicy_FrameOptionsProperty
	// Experimental.
	InternalValue() *AwsResponseHeadersPolicy_SecurityHeadersConfigProperty
	// Experimental.
	SetInternalValue(val *AwsResponseHeadersPolicy_SecurityHeadersConfigProperty)
	// Experimental.
	ReferrerPolicy() AwsResponseHeadersPolicy_ReferrerPolicyPropertyOutputReference
	// Experimental.
	ReferrerPolicyInput() *AwsResponseHeadersPolicy_ReferrerPolicyProperty
	// Experimental.
	StrictTransportSecurity() AwsResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference
	// Experimental.
	StrictTransportSecurityInput() *AwsResponseHeadersPolicy_StrictTransportSecurityProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	XssProtection() AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference
	// Experimental.
	XssProtectionInput() *AwsResponseHeadersPolicy_XssProtectionProperty
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
	PutContentSecurityPolicy(value *AwsResponseHeadersPolicy_ContentSecurityPolicyProperty)
	// Experimental.
	PutContentTypeOptions(value *AwsResponseHeadersPolicy_ContentTypeOptionsProperty)
	// Experimental.
	PutFrameOptions(value *AwsResponseHeadersPolicy_FrameOptionsProperty)
	// Experimental.
	PutReferrerPolicy(value *AwsResponseHeadersPolicy_ReferrerPolicyProperty)
	// Experimental.
	PutStrictTransportSecurity(value *AwsResponseHeadersPolicy_StrictTransportSecurityProperty)
	// Experimental.
	PutXssProtection(value *AwsResponseHeadersPolicy_XssProtectionProperty)
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

// The jsii proxy struct for AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference
type jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ContentSecurityPolicy() AwsResponseHeadersPolicy_ContentSecurityPolicyPropertyOutputReference {
	var returns AwsResponseHeadersPolicy_ContentSecurityPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"contentSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ContentSecurityPolicyInput() *AwsResponseHeadersPolicy_ContentSecurityPolicyProperty {
	var returns *AwsResponseHeadersPolicy_ContentSecurityPolicyProperty
	_jsii_.Get(
		j,
		"contentSecurityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ContentTypeOptions() AwsResponseHeadersPolicy_ContentTypeOptionsPropertyOutputReference {
	var returns AwsResponseHeadersPolicy_ContentTypeOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"contentTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ContentTypeOptionsInput() *AwsResponseHeadersPolicy_ContentTypeOptionsProperty {
	var returns *AwsResponseHeadersPolicy_ContentTypeOptionsProperty
	_jsii_.Get(
		j,
		"contentTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) FrameOptions() AwsResponseHeadersPolicy_FrameOptionsPropertyOutputReference {
	var returns AwsResponseHeadersPolicy_FrameOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"frameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) FrameOptionsInput() *AwsResponseHeadersPolicy_FrameOptionsProperty {
	var returns *AwsResponseHeadersPolicy_FrameOptionsProperty
	_jsii_.Get(
		j,
		"frameOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) InternalValue() *AwsResponseHeadersPolicy_SecurityHeadersConfigProperty {
	var returns *AwsResponseHeadersPolicy_SecurityHeadersConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ReferrerPolicy() AwsResponseHeadersPolicy_ReferrerPolicyPropertyOutputReference {
	var returns AwsResponseHeadersPolicy_ReferrerPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"referrerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ReferrerPolicyInput() *AwsResponseHeadersPolicy_ReferrerPolicyProperty {
	var returns *AwsResponseHeadersPolicy_ReferrerPolicyProperty
	_jsii_.Get(
		j,
		"referrerPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) StrictTransportSecurity() AwsResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference {
	var returns AwsResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference
	_jsii_.Get(
		j,
		"strictTransportSecurity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) StrictTransportSecurityInput() *AwsResponseHeadersPolicy_StrictTransportSecurityProperty {
	var returns *AwsResponseHeadersPolicy_StrictTransportSecurityProperty
	_jsii_.Get(
		j,
		"strictTransportSecurityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) XssProtection() AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference {
	var returns AwsResponseHeadersPolicy_XssProtectionPropertyOutputReference
	_jsii_.Get(
		j,
		"xssProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) XssProtectionInput() *AwsResponseHeadersPolicy_XssProtectionProperty {
	var returns *AwsResponseHeadersPolicy_XssProtectionProperty
	_jsii_.Get(
		j,
		"xssProtectionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsResponseHeadersPolicy.SecurityHeadersConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference_Override(a AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsResponseHeadersPolicy.SecurityHeadersConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference)SetInternalValue(val *AwsResponseHeadersPolicy_SecurityHeadersConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutContentSecurityPolicy(value *AwsResponseHeadersPolicy_ContentSecurityPolicyProperty) {
	if err := a.validatePutContentSecurityPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContentSecurityPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutContentTypeOptions(value *AwsResponseHeadersPolicy_ContentTypeOptionsProperty) {
	if err := a.validatePutContentTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContentTypeOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutFrameOptions(value *AwsResponseHeadersPolicy_FrameOptionsProperty) {
	if err := a.validatePutFrameOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFrameOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutReferrerPolicy(value *AwsResponseHeadersPolicy_ReferrerPolicyProperty) {
	if err := a.validatePutReferrerPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReferrerPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutStrictTransportSecurity(value *AwsResponseHeadersPolicy_StrictTransportSecurityProperty) {
	if err := a.validatePutStrictTransportSecurityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStrictTransportSecurity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutXssProtection(value *AwsResponseHeadersPolicy_XssProtectionProperty) {
	if err := a.validatePutXssProtectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putXssProtection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetContentSecurityPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetContentSecurityPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetContentTypeOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetContentTypeOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetFrameOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetFrameOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetReferrerPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetReferrerPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetStrictTransportSecurity() {
	_jsii_.InvokeVoid(
		a,
		"resetStrictTransportSecurity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetXssProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetXssProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

