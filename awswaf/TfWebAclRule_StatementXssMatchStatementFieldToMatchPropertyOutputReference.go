package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllQueryArguments() TfWebAclRule_StatementXssMatchStatementFieldToMatchAllQueryArgumentsPropertyList
	// Experimental.
	AllQueryArgumentsInput() interface{}
	// Experimental.
	Body() TfWebAclRule_StatementXssMatchStatementFieldToMatchBodyPropertyList
	// Experimental.
	BodyInput() interface{}
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
	Cookies() TfWebAclRule_StatementXssMatchStatementFieldToMatchCookiesPropertyList
	// Experimental.
	CookiesInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	HeaderOrder() TfWebAclRule_StatementXssMatchStatementFieldToMatchHeaderOrderPropertyList
	// Experimental.
	HeaderOrderInput() interface{}
	// Experimental.
	Headers() TfWebAclRule_StatementXssMatchStatementFieldToMatchHeadersPropertyList
	// Experimental.
	HeadersInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Ja3Fingerprint() TfWebAclRule_StatementXssMatchStatementFieldToMatchJa3FingerprintPropertyList
	// Experimental.
	Ja3FingerprintInput() interface{}
	// Experimental.
	Ja4Fingerprint() TfWebAclRule_StatementXssMatchStatementFieldToMatchJa4FingerprintPropertyList
	// Experimental.
	Ja4FingerprintInput() interface{}
	// Experimental.
	JsonBody() TfWebAclRule_StatementXssMatchStatementFieldToMatchJsonBodyPropertyList
	// Experimental.
	JsonBodyInput() interface{}
	// Experimental.
	Method() TfWebAclRule_StatementXssMatchStatementFieldToMatchMethodPropertyList
	// Experimental.
	MethodInput() interface{}
	// Experimental.
	QueryString() TfWebAclRule_StatementXssMatchStatementFieldToMatchQueryStringPropertyList
	// Experimental.
	QueryStringInput() interface{}
	// Experimental.
	SingleHeader() TfWebAclRule_StatementXssMatchStatementFieldToMatchSingleHeaderPropertyList
	// Experimental.
	SingleHeaderInput() interface{}
	// Experimental.
	SingleQueryArgument() TfWebAclRule_StatementXssMatchStatementFieldToMatchSingleQueryArgumentPropertyList
	// Experimental.
	SingleQueryArgumentInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UriFragment() TfWebAclRule_StatementXssMatchStatementFieldToMatchUriFragmentPropertyList
	// Experimental.
	UriFragmentInput() interface{}
	// Experimental.
	UriPath() TfWebAclRule_StatementXssMatchStatementFieldToMatchUriPathPropertyList
	// Experimental.
	UriPathInput() interface{}
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
	PutAllQueryArguments(value interface{})
	// Experimental.
	PutBody(value interface{})
	// Experimental.
	PutCookies(value interface{})
	// Experimental.
	PutHeaderOrder(value interface{})
	// Experimental.
	PutHeaders(value interface{})
	// Experimental.
	PutJa3Fingerprint(value interface{})
	// Experimental.
	PutJa4Fingerprint(value interface{})
	// Experimental.
	PutJsonBody(value interface{})
	// Experimental.
	PutMethod(value interface{})
	// Experimental.
	PutQueryString(value interface{})
	// Experimental.
	PutSingleHeader(value interface{})
	// Experimental.
	PutSingleQueryArgument(value interface{})
	// Experimental.
	PutUriFragment(value interface{})
	// Experimental.
	PutUriPath(value interface{})
	// Experimental.
	ResetAllQueryArguments()
	// Experimental.
	ResetBody()
	// Experimental.
	ResetCookies()
	// Experimental.
	ResetHeaderOrder()
	// Experimental.
	ResetHeaders()
	// Experimental.
	ResetJa3Fingerprint()
	// Experimental.
	ResetJa4Fingerprint()
	// Experimental.
	ResetJsonBody()
	// Experimental.
	ResetMethod()
	// Experimental.
	ResetQueryString()
	// Experimental.
	ResetSingleHeader()
	// Experimental.
	ResetSingleQueryArgument()
	// Experimental.
	ResetUriFragment()
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

// The jsii proxy struct for TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference
type jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) AllQueryArguments() TfWebAclRule_StatementXssMatchStatementFieldToMatchAllQueryArgumentsPropertyList {
	var returns TfWebAclRule_StatementXssMatchStatementFieldToMatchAllQueryArgumentsPropertyList
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) AllQueryArgumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) Body() TfWebAclRule_StatementXssMatchStatementFieldToMatchBodyPropertyList {
	var returns TfWebAclRule_StatementXssMatchStatementFieldToMatchBodyPropertyList
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) BodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) Cookies() TfWebAclRule_StatementXssMatchStatementFieldToMatchCookiesPropertyList {
	var returns TfWebAclRule_StatementXssMatchStatementFieldToMatchCookiesPropertyList
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) CookiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) HeaderOrder() TfWebAclRule_StatementXssMatchStatementFieldToMatchHeaderOrderPropertyList {
	var returns TfWebAclRule_StatementXssMatchStatementFieldToMatchHeaderOrderPropertyList
	_jsii_.Get(
		j,
		"headerOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) HeaderOrderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) Headers() TfWebAclRule_StatementXssMatchStatementFieldToMatchHeadersPropertyList {
	var returns TfWebAclRule_StatementXssMatchStatementFieldToMatchHeadersPropertyList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) Ja3Fingerprint() TfWebAclRule_StatementXssMatchStatementFieldToMatchJa3FingerprintPropertyList {
	var returns TfWebAclRule_StatementXssMatchStatementFieldToMatchJa3FingerprintPropertyList
	_jsii_.Get(
		j,
		"ja3Fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) Ja3FingerprintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ja3FingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) Ja4Fingerprint() TfWebAclRule_StatementXssMatchStatementFieldToMatchJa4FingerprintPropertyList {
	var returns TfWebAclRule_StatementXssMatchStatementFieldToMatchJa4FingerprintPropertyList
	_jsii_.Get(
		j,
		"ja4Fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) Ja4FingerprintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ja4FingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) JsonBody() TfWebAclRule_StatementXssMatchStatementFieldToMatchJsonBodyPropertyList {
	var returns TfWebAclRule_StatementXssMatchStatementFieldToMatchJsonBodyPropertyList
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) JsonBodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) Method() TfWebAclRule_StatementXssMatchStatementFieldToMatchMethodPropertyList {
	var returns TfWebAclRule_StatementXssMatchStatementFieldToMatchMethodPropertyList
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) MethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) QueryString() TfWebAclRule_StatementXssMatchStatementFieldToMatchQueryStringPropertyList {
	var returns TfWebAclRule_StatementXssMatchStatementFieldToMatchQueryStringPropertyList
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) SingleHeader() TfWebAclRule_StatementXssMatchStatementFieldToMatchSingleHeaderPropertyList {
	var returns TfWebAclRule_StatementXssMatchStatementFieldToMatchSingleHeaderPropertyList
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) SingleHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) SingleQueryArgument() TfWebAclRule_StatementXssMatchStatementFieldToMatchSingleQueryArgumentPropertyList {
	var returns TfWebAclRule_StatementXssMatchStatementFieldToMatchSingleQueryArgumentPropertyList
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) SingleQueryArgumentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) UriFragment() TfWebAclRule_StatementXssMatchStatementFieldToMatchUriFragmentPropertyList {
	var returns TfWebAclRule_StatementXssMatchStatementFieldToMatchUriFragmentPropertyList
	_jsii_.Get(
		j,
		"uriFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) UriFragmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uriFragmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) UriPath() TfWebAclRule_StatementXssMatchStatementFieldToMatchUriPathPropertyList {
	var returns TfWebAclRule_StatementXssMatchStatementFieldToMatchUriPathPropertyList
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) UriPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementXssMatchStatementFieldToMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference_Override(t TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementXssMatchStatementFieldToMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) PutAllQueryArguments(value interface{}) {
	if err := t.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) PutBody(value interface{}) {
	if err := t.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBody",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) PutCookies(value interface{}) {
	if err := t.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCookies",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) PutHeaderOrder(value interface{}) {
	if err := t.validatePutHeaderOrderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHeaderOrder",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) PutHeaders(value interface{}) {
	if err := t.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHeaders",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) PutJa3Fingerprint(value interface{}) {
	if err := t.validatePutJa3FingerprintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJa3Fingerprint",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) PutJa4Fingerprint(value interface{}) {
	if err := t.validatePutJa4FingerprintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJa4Fingerprint",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) PutJsonBody(value interface{}) {
	if err := t.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) PutMethod(value interface{}) {
	if err := t.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMethod",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) PutQueryString(value interface{}) {
	if err := t.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putQueryString",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) PutSingleHeader(value interface{}) {
	if err := t.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) PutSingleQueryArgument(value interface{}) {
	if err := t.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) PutUriFragment(value interface{}) {
	if err := t.validatePutUriFragmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUriFragment",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) PutUriPath(value interface{}) {
	if err := t.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUriPath",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		t,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		t,
		"resetBody",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		t,
		"resetCookies",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetHeaderOrder() {
	_jsii_.InvokeVoid(
		t,
		"resetHeaderOrder",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		t,
		"resetHeaders",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetJa3Fingerprint() {
	_jsii_.InvokeVoid(
		t,
		"resetJa3Fingerprint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetJa4Fingerprint() {
	_jsii_.InvokeVoid(
		t,
		"resetJa4Fingerprint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		t,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		t,
		"resetMethod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		t,
		"resetQueryString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		t,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetUriFragment() {
	_jsii_.InvokeVoid(
		t,
		"resetUriFragment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		t,
		"resetUriPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementXssMatchStatementFieldToMatchPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

