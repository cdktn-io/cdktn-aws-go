package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfResponseHeadersPolicy_CorsConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessControlAllowCredentials() interface{}
	// Experimental.
	SetAccessControlAllowCredentials(val interface{})
	// Experimental.
	AccessControlAllowCredentialsInput() interface{}
	// Experimental.
	AccessControlAllowHeaders() TfResponseHeadersPolicy_AccessControlAllowHeadersPropertyOutputReference
	// Experimental.
	AccessControlAllowHeadersInput() *TfResponseHeadersPolicy_AccessControlAllowHeadersProperty
	// Experimental.
	AccessControlAllowMethods() TfResponseHeadersPolicy_AccessControlAllowMethodsPropertyOutputReference
	// Experimental.
	AccessControlAllowMethodsInput() *TfResponseHeadersPolicy_AccessControlAllowMethodsProperty
	// Experimental.
	AccessControlAllowOrigins() TfResponseHeadersPolicy_AccessControlAllowOriginsPropertyOutputReference
	// Experimental.
	AccessControlAllowOriginsInput() *TfResponseHeadersPolicy_AccessControlAllowOriginsProperty
	// Experimental.
	AccessControlExposeHeaders() TfResponseHeadersPolicy_AccessControlExposeHeadersPropertyOutputReference
	// Experimental.
	AccessControlExposeHeadersInput() *TfResponseHeadersPolicy_AccessControlExposeHeadersProperty
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
	InternalValue() *TfResponseHeadersPolicy_CorsConfigProperty
	// Experimental.
	SetInternalValue(val *TfResponseHeadersPolicy_CorsConfigProperty)
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
	PutAccessControlAllowHeaders(value *TfResponseHeadersPolicy_AccessControlAllowHeadersProperty)
	// Experimental.
	PutAccessControlAllowMethods(value *TfResponseHeadersPolicy_AccessControlAllowMethodsProperty)
	// Experimental.
	PutAccessControlAllowOrigins(value *TfResponseHeadersPolicy_AccessControlAllowOriginsProperty)
	// Experimental.
	PutAccessControlExposeHeaders(value *TfResponseHeadersPolicy_AccessControlExposeHeadersProperty)
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

// The jsii proxy struct for TfResponseHeadersPolicy_CorsConfigPropertyOutputReference
type jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlAllowCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlAllowCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlAllowCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlAllowCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlAllowHeaders() TfResponseHeadersPolicy_AccessControlAllowHeadersPropertyOutputReference {
	var returns TfResponseHeadersPolicy_AccessControlAllowHeadersPropertyOutputReference
	_jsii_.Get(
		j,
		"accessControlAllowHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlAllowHeadersInput() *TfResponseHeadersPolicy_AccessControlAllowHeadersProperty {
	var returns *TfResponseHeadersPolicy_AccessControlAllowHeadersProperty
	_jsii_.Get(
		j,
		"accessControlAllowHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlAllowMethods() TfResponseHeadersPolicy_AccessControlAllowMethodsPropertyOutputReference {
	var returns TfResponseHeadersPolicy_AccessControlAllowMethodsPropertyOutputReference
	_jsii_.Get(
		j,
		"accessControlAllowMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlAllowMethodsInput() *TfResponseHeadersPolicy_AccessControlAllowMethodsProperty {
	var returns *TfResponseHeadersPolicy_AccessControlAllowMethodsProperty
	_jsii_.Get(
		j,
		"accessControlAllowMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlAllowOrigins() TfResponseHeadersPolicy_AccessControlAllowOriginsPropertyOutputReference {
	var returns TfResponseHeadersPolicy_AccessControlAllowOriginsPropertyOutputReference
	_jsii_.Get(
		j,
		"accessControlAllowOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlAllowOriginsInput() *TfResponseHeadersPolicy_AccessControlAllowOriginsProperty {
	var returns *TfResponseHeadersPolicy_AccessControlAllowOriginsProperty
	_jsii_.Get(
		j,
		"accessControlAllowOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlExposeHeaders() TfResponseHeadersPolicy_AccessControlExposeHeadersPropertyOutputReference {
	var returns TfResponseHeadersPolicy_AccessControlExposeHeadersPropertyOutputReference
	_jsii_.Get(
		j,
		"accessControlExposeHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlExposeHeadersInput() *TfResponseHeadersPolicy_AccessControlExposeHeadersProperty {
	var returns *TfResponseHeadersPolicy_AccessControlExposeHeadersProperty
	_jsii_.Get(
		j,
		"accessControlExposeHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlMaxAgeSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) AccessControlMaxAgeSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) InternalValue() *TfResponseHeadersPolicy_CorsConfigProperty {
	var returns *TfResponseHeadersPolicy_CorsConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) OriginOverride() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) OriginOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfResponseHeadersPolicy_CorsConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfResponseHeadersPolicy_CorsConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfResponseHeadersPolicy_CorsConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfResponseHeadersPolicy.CorsConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfResponseHeadersPolicy_CorsConfigPropertyOutputReference_Override(t TfResponseHeadersPolicy_CorsConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfResponseHeadersPolicy.CorsConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference)SetAccessControlAllowCredentials(val interface{}) {
	if err := j.validateSetAccessControlAllowCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessControlAllowCredentials",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference)SetAccessControlMaxAgeSec(val *float64) {
	if err := j.validateSetAccessControlMaxAgeSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessControlMaxAgeSec",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference)SetInternalValue(val *TfResponseHeadersPolicy_CorsConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference)SetOriginOverride(val interface{}) {
	if err := j.validateSetOriginOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originOverride",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) PutAccessControlAllowHeaders(value *TfResponseHeadersPolicy_AccessControlAllowHeadersProperty) {
	if err := t.validatePutAccessControlAllowHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAccessControlAllowHeaders",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) PutAccessControlAllowMethods(value *TfResponseHeadersPolicy_AccessControlAllowMethodsProperty) {
	if err := t.validatePutAccessControlAllowMethodsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAccessControlAllowMethods",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) PutAccessControlAllowOrigins(value *TfResponseHeadersPolicy_AccessControlAllowOriginsProperty) {
	if err := t.validatePutAccessControlAllowOriginsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAccessControlAllowOrigins",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) PutAccessControlExposeHeaders(value *TfResponseHeadersPolicy_AccessControlExposeHeadersProperty) {
	if err := t.validatePutAccessControlExposeHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAccessControlExposeHeaders",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) ResetAccessControlExposeHeaders() {
	_jsii_.InvokeVoid(
		t,
		"resetAccessControlExposeHeaders",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) ResetAccessControlMaxAgeSec() {
	_jsii_.InvokeVoid(
		t,
		"resetAccessControlMaxAgeSec",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfResponseHeadersPolicy_CorsConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

