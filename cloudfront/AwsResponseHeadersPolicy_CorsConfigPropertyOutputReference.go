package cloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessControlAllowCredentials() interface{}
	// Experimental.
	SetAccessControlAllowCredentials(val interface{})
	// Experimental.
	AccessControlAllowCredentialsInput() interface{}
	// Experimental.
	AccessControlAllowHeaders() AwsResponseHeadersPolicy_AccessControlAllowHeadersPropertyOutputReference
	// Experimental.
	AccessControlAllowHeadersInput() *AwsResponseHeadersPolicy_AccessControlAllowHeadersProperty
	// Experimental.
	AccessControlAllowMethods() AwsResponseHeadersPolicy_AccessControlAllowMethodsPropertyOutputReference
	// Experimental.
	AccessControlAllowMethodsInput() *AwsResponseHeadersPolicy_AccessControlAllowMethodsProperty
	// Experimental.
	AccessControlAllowOrigins() AwsResponseHeadersPolicy_AccessControlAllowOriginsPropertyOutputReference
	// Experimental.
	AccessControlAllowOriginsInput() *AwsResponseHeadersPolicy_AccessControlAllowOriginsProperty
	// Experimental.
	AccessControlExposeHeaders() AwsResponseHeadersPolicy_AccessControlExposeHeadersPropertyOutputReference
	// Experimental.
	AccessControlExposeHeadersInput() *AwsResponseHeadersPolicy_AccessControlExposeHeadersProperty
	// Experimental.
	AccessControlMaxAgeSec() *float64
	// Experimental.
	SetAccessControlMaxAgeSec(val *float64)
	// Experimental.
	AccessControlMaxAgeSecInput() *float64
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
	InternalValue() *AwsResponseHeadersPolicy_CorsConfigProperty
	// Experimental.
	SetInternalValue(val *AwsResponseHeadersPolicy_CorsConfigProperty)
	// Experimental.
	OriginOverride() interface{}
	// Experimental.
	SetOriginOverride(val interface{})
	// Experimental.
	OriginOverrideInput() interface{}
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
	PutAccessControlAllowHeaders(value *AwsResponseHeadersPolicy_AccessControlAllowHeadersProperty)
	// Experimental.
	PutAccessControlAllowMethods(value *AwsResponseHeadersPolicy_AccessControlAllowMethodsProperty)
	// Experimental.
	PutAccessControlAllowOrigins(value *AwsResponseHeadersPolicy_AccessControlAllowOriginsProperty)
	// Experimental.
	PutAccessControlExposeHeaders(value *AwsResponseHeadersPolicy_AccessControlExposeHeadersProperty)
	// Experimental.
	ResetAccessControlExposeHeaders()
	// Experimental.
	ResetAccessControlMaxAgeSec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference
type jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlAllowCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlAllowCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlAllowCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlAllowCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlAllowHeaders() AwsResponseHeadersPolicy_AccessControlAllowHeadersPropertyOutputReference {
	var returns AwsResponseHeadersPolicy_AccessControlAllowHeadersPropertyOutputReference
	_jsii_.Get(
		j,
		"accessControlAllowHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlAllowHeadersInput() *AwsResponseHeadersPolicy_AccessControlAllowHeadersProperty {
	var returns *AwsResponseHeadersPolicy_AccessControlAllowHeadersProperty
	_jsii_.Get(
		j,
		"accessControlAllowHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlAllowMethods() AwsResponseHeadersPolicy_AccessControlAllowMethodsPropertyOutputReference {
	var returns AwsResponseHeadersPolicy_AccessControlAllowMethodsPropertyOutputReference
	_jsii_.Get(
		j,
		"accessControlAllowMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlAllowMethodsInput() *AwsResponseHeadersPolicy_AccessControlAllowMethodsProperty {
	var returns *AwsResponseHeadersPolicy_AccessControlAllowMethodsProperty
	_jsii_.Get(
		j,
		"accessControlAllowMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlAllowOrigins() AwsResponseHeadersPolicy_AccessControlAllowOriginsPropertyOutputReference {
	var returns AwsResponseHeadersPolicy_AccessControlAllowOriginsPropertyOutputReference
	_jsii_.Get(
		j,
		"accessControlAllowOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlAllowOriginsInput() *AwsResponseHeadersPolicy_AccessControlAllowOriginsProperty {
	var returns *AwsResponseHeadersPolicy_AccessControlAllowOriginsProperty
	_jsii_.Get(
		j,
		"accessControlAllowOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlExposeHeaders() AwsResponseHeadersPolicy_AccessControlExposeHeadersPropertyOutputReference {
	var returns AwsResponseHeadersPolicy_AccessControlExposeHeadersPropertyOutputReference
	_jsii_.Get(
		j,
		"accessControlExposeHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlExposeHeadersInput() *AwsResponseHeadersPolicy_AccessControlExposeHeadersProperty {
	var returns *AwsResponseHeadersPolicy_AccessControlExposeHeadersProperty
	_jsii_.Get(
		j,
		"accessControlExposeHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlMaxAgeSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlMaxAgeSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) InternalValue() *AwsResponseHeadersPolicy_CorsConfigProperty {
	var returns *AwsResponseHeadersPolicy_CorsConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) OriginOverride() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) OriginOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsResponseHeadersPolicy_CorsConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsResponseHeadersPolicy_CorsConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsResponseHeadersPolicy.CorsConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsResponseHeadersPolicy_CorsConfigPropertyOutputReference_Override(a AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsResponseHeadersPolicy.CorsConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference)SetAccessControlAllowCredentials(val interface{}) {
	if err := j.validateSetAccessControlAllowCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessControlAllowCredentials",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference)SetAccessControlMaxAgeSec(val *float64) {
	if err := j.validateSetAccessControlMaxAgeSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessControlMaxAgeSec",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference)SetInternalValue(val *AwsResponseHeadersPolicy_CorsConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference)SetOriginOverride(val interface{}) {
	if err := j.validateSetOriginOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originOverride",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) PutAccessControlAllowHeaders(value *AwsResponseHeadersPolicy_AccessControlAllowHeadersProperty) {
	if err := a.validatePutAccessControlAllowHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessControlAllowHeaders",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) PutAccessControlAllowMethods(value *AwsResponseHeadersPolicy_AccessControlAllowMethodsProperty) {
	if err := a.validatePutAccessControlAllowMethodsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessControlAllowMethods",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) PutAccessControlAllowOrigins(value *AwsResponseHeadersPolicy_AccessControlAllowOriginsProperty) {
	if err := a.validatePutAccessControlAllowOriginsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessControlAllowOrigins",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) PutAccessControlExposeHeaders(value *AwsResponseHeadersPolicy_AccessControlExposeHeadersProperty) {
	if err := a.validatePutAccessControlExposeHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessControlExposeHeaders",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) ResetAccessControlExposeHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessControlExposeHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) ResetAccessControlMaxAgeSec() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessControlMaxAgeSec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsResponseHeadersPolicy_CorsConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

