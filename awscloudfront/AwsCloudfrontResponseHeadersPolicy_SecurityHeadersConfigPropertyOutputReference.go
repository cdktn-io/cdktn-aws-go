package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference interface {
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
	ContentSecurityPolicy() AwsCloudfrontResponseHeadersPolicy_ContentSecurityPolicyPropertyOutputReference
	// Experimental.
	ContentSecurityPolicyInput() *AwsCloudfrontResponseHeadersPolicy_ContentSecurityPolicyProperty
	// Experimental.
	ContentTypeOptions() AwsCloudfrontResponseHeadersPolicy_ContentTypeOptionsPropertyOutputReference
	// Experimental.
	ContentTypeOptionsInput() *AwsCloudfrontResponseHeadersPolicy_ContentTypeOptionsProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	FrameOptions() AwsCloudfrontResponseHeadersPolicy_FrameOptionsPropertyOutputReference
	// Experimental.
	FrameOptionsInput() *AwsCloudfrontResponseHeadersPolicy_FrameOptionsProperty
	// Experimental.
	InternalValue() *AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigProperty
	// Experimental.
	SetInternalValue(val *AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigProperty)
	// Experimental.
	ReferrerPolicy() AwsCloudfrontResponseHeadersPolicy_ReferrerPolicyPropertyOutputReference
	// Experimental.
	ReferrerPolicyInput() *AwsCloudfrontResponseHeadersPolicy_ReferrerPolicyProperty
	// Experimental.
	StrictTransportSecurity() AwsCloudfrontResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference
	// Experimental.
	StrictTransportSecurityInput() *AwsCloudfrontResponseHeadersPolicy_StrictTransportSecurityProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	XssProtection() AwsCloudfrontResponseHeadersPolicy_XssProtectionPropertyOutputReference
	// Experimental.
	XssProtectionInput() *AwsCloudfrontResponseHeadersPolicy_XssProtectionProperty
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
	PutContentSecurityPolicy(value *AwsCloudfrontResponseHeadersPolicy_ContentSecurityPolicyProperty)
	// Experimental.
	PutContentTypeOptions(value *AwsCloudfrontResponseHeadersPolicy_ContentTypeOptionsProperty)
	// Experimental.
	PutFrameOptions(value *AwsCloudfrontResponseHeadersPolicy_FrameOptionsProperty)
	// Experimental.
	PutReferrerPolicy(value *AwsCloudfrontResponseHeadersPolicy_ReferrerPolicyProperty)
	// Experimental.
	PutStrictTransportSecurity(value *AwsCloudfrontResponseHeadersPolicy_StrictTransportSecurityProperty)
	// Experimental.
	PutXssProtection(value *AwsCloudfrontResponseHeadersPolicy_XssProtectionProperty)
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

// The jsii proxy struct for AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference
type jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ContentSecurityPolicy() AwsCloudfrontResponseHeadersPolicy_ContentSecurityPolicyPropertyOutputReference {
	var returns AwsCloudfrontResponseHeadersPolicy_ContentSecurityPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"contentSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ContentSecurityPolicyInput() *AwsCloudfrontResponseHeadersPolicy_ContentSecurityPolicyProperty {
	var returns *AwsCloudfrontResponseHeadersPolicy_ContentSecurityPolicyProperty
	_jsii_.Get(
		j,
		"contentSecurityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ContentTypeOptions() AwsCloudfrontResponseHeadersPolicy_ContentTypeOptionsPropertyOutputReference {
	var returns AwsCloudfrontResponseHeadersPolicy_ContentTypeOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"contentTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ContentTypeOptionsInput() *AwsCloudfrontResponseHeadersPolicy_ContentTypeOptionsProperty {
	var returns *AwsCloudfrontResponseHeadersPolicy_ContentTypeOptionsProperty
	_jsii_.Get(
		j,
		"contentTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) FrameOptions() AwsCloudfrontResponseHeadersPolicy_FrameOptionsPropertyOutputReference {
	var returns AwsCloudfrontResponseHeadersPolicy_FrameOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"frameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) FrameOptionsInput() *AwsCloudfrontResponseHeadersPolicy_FrameOptionsProperty {
	var returns *AwsCloudfrontResponseHeadersPolicy_FrameOptionsProperty
	_jsii_.Get(
		j,
		"frameOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) InternalValue() *AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigProperty {
	var returns *AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ReferrerPolicy() AwsCloudfrontResponseHeadersPolicy_ReferrerPolicyPropertyOutputReference {
	var returns AwsCloudfrontResponseHeadersPolicy_ReferrerPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"referrerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ReferrerPolicyInput() *AwsCloudfrontResponseHeadersPolicy_ReferrerPolicyProperty {
	var returns *AwsCloudfrontResponseHeadersPolicy_ReferrerPolicyProperty
	_jsii_.Get(
		j,
		"referrerPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) StrictTransportSecurity() AwsCloudfrontResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference {
	var returns AwsCloudfrontResponseHeadersPolicy_StrictTransportSecurityPropertyOutputReference
	_jsii_.Get(
		j,
		"strictTransportSecurity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) StrictTransportSecurityInput() *AwsCloudfrontResponseHeadersPolicy_StrictTransportSecurityProperty {
	var returns *AwsCloudfrontResponseHeadersPolicy_StrictTransportSecurityProperty
	_jsii_.Get(
		j,
		"strictTransportSecurityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) XssProtection() AwsCloudfrontResponseHeadersPolicy_XssProtectionPropertyOutputReference {
	var returns AwsCloudfrontResponseHeadersPolicy_XssProtectionPropertyOutputReference
	_jsii_.Get(
		j,
		"xssProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) XssProtectionInput() *AwsCloudfrontResponseHeadersPolicy_XssProtectionProperty {
	var returns *AwsCloudfrontResponseHeadersPolicy_XssProtectionProperty
	_jsii_.Get(
		j,
		"xssProtectionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsCloudfrontResponseHeadersPolicy.SecurityHeadersConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference_Override(a AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsCloudfrontResponseHeadersPolicy.SecurityHeadersConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference)SetInternalValue(val *AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutContentSecurityPolicy(value *AwsCloudfrontResponseHeadersPolicy_ContentSecurityPolicyProperty) {
	if err := a.validatePutContentSecurityPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContentSecurityPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutContentTypeOptions(value *AwsCloudfrontResponseHeadersPolicy_ContentTypeOptionsProperty) {
	if err := a.validatePutContentTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContentTypeOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutFrameOptions(value *AwsCloudfrontResponseHeadersPolicy_FrameOptionsProperty) {
	if err := a.validatePutFrameOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFrameOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutReferrerPolicy(value *AwsCloudfrontResponseHeadersPolicy_ReferrerPolicyProperty) {
	if err := a.validatePutReferrerPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReferrerPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutStrictTransportSecurity(value *AwsCloudfrontResponseHeadersPolicy_StrictTransportSecurityProperty) {
	if err := a.validatePutStrictTransportSecurityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStrictTransportSecurity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) PutXssProtection(value *AwsCloudfrontResponseHeadersPolicy_XssProtectionProperty) {
	if err := a.validatePutXssProtectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putXssProtection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetContentSecurityPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetContentSecurityPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetContentTypeOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetContentTypeOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetFrameOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetFrameOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetReferrerPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetReferrerPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetStrictTransportSecurity() {
	_jsii_.InvokeVoid(
		a,
		"resetStrictTransportSecurity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ResetXssProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetXssProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

