package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Method() TfWebAclLoggingConfiguration_MethodPropertyOutputReference
	// Experimental.
	MethodInput() *TfWebAclLoggingConfiguration_MethodProperty
	// Experimental.
	QueryString() TfWebAclLoggingConfiguration_QueryStringPropertyOutputReference
	// Experimental.
	QueryStringInput() *TfWebAclLoggingConfiguration_QueryStringProperty
	// Experimental.
	SingleHeader() TfWebAclLoggingConfiguration_SingleHeaderPropertyOutputReference
	// Experimental.
	SingleHeaderInput() *TfWebAclLoggingConfiguration_SingleHeaderProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UriPath() TfWebAclLoggingConfiguration_UriPathPropertyOutputReference
	// Experimental.
	UriPathInput() *TfWebAclLoggingConfiguration_UriPathProperty
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
	PutMethod(value *TfWebAclLoggingConfiguration_MethodProperty)
	// Experimental.
	PutQueryString(value *TfWebAclLoggingConfiguration_QueryStringProperty)
	// Experimental.
	PutSingleHeader(value *TfWebAclLoggingConfiguration_SingleHeaderProperty)
	// Experimental.
	PutUriPath(value *TfWebAclLoggingConfiguration_UriPathProperty)
	// Experimental.
	ResetMethod()
	// Experimental.
	ResetQueryString()
	// Experimental.
	ResetSingleHeader()
	// Experimental.
	ResetUriPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference
type jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) Method() TfWebAclLoggingConfiguration_MethodPropertyOutputReference {
	var returns TfWebAclLoggingConfiguration_MethodPropertyOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) MethodInput() *TfWebAclLoggingConfiguration_MethodProperty {
	var returns *TfWebAclLoggingConfiguration_MethodProperty
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) QueryString() TfWebAclLoggingConfiguration_QueryStringPropertyOutputReference {
	var returns TfWebAclLoggingConfiguration_QueryStringPropertyOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) QueryStringInput() *TfWebAclLoggingConfiguration_QueryStringProperty {
	var returns *TfWebAclLoggingConfiguration_QueryStringProperty
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) SingleHeader() TfWebAclLoggingConfiguration_SingleHeaderPropertyOutputReference {
	var returns TfWebAclLoggingConfiguration_SingleHeaderPropertyOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) SingleHeaderInput() *TfWebAclLoggingConfiguration_SingleHeaderProperty {
	var returns *TfWebAclLoggingConfiguration_SingleHeaderProperty
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) UriPath() TfWebAclLoggingConfiguration_UriPathPropertyOutputReference {
	var returns TfWebAclLoggingConfiguration_UriPathPropertyOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) UriPathInput() *TfWebAclLoggingConfiguration_UriPathProperty {
	var returns *TfWebAclLoggingConfiguration_UriPathProperty
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclLoggingConfiguration.RedactedFieldsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference_Override(t TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclLoggingConfiguration.RedactedFieldsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) PutMethod(value *TfWebAclLoggingConfiguration_MethodProperty) {
	if err := t.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMethod",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) PutQueryString(value *TfWebAclLoggingConfiguration_QueryStringProperty) {
	if err := t.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putQueryString",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) PutSingleHeader(value *TfWebAclLoggingConfiguration_SingleHeaderProperty) {
	if err := t.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) PutUriPath(value *TfWebAclLoggingConfiguration_UriPathProperty) {
	if err := t.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUriPath",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		t,
		"resetMethod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		t,
		"resetQueryString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		t,
		"resetUriPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclLoggingConfiguration_RedactedFieldsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

