package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllQueryArguments() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsPropertyList
	// Experimental.
	AllQueryArgumentsInput() interface{}
	// Experimental.
	Body() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchBodyPropertyList
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
	Cookies() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchCookiesPropertyList
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
	HeaderOrder() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchHeaderOrderPropertyList
	// Experimental.
	HeaderOrderInput() interface{}
	// Experimental.
	Headers() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchHeadersPropertyList
	// Experimental.
	HeadersInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Ja3Fingerprint() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJa3FingerprintPropertyList
	// Experimental.
	Ja3FingerprintInput() interface{}
	// Experimental.
	Ja4Fingerprint() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJa4FingerprintPropertyList
	// Experimental.
	Ja4FingerprintInput() interface{}
	// Experimental.
	JsonBody() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyPropertyList
	// Experimental.
	JsonBodyInput() interface{}
	// Experimental.
	Method() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList
	// Experimental.
	MethodInput() interface{}
	// Experimental.
	QueryString() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchQueryStringPropertyList
	// Experimental.
	QueryStringInput() interface{}
	// Experimental.
	SingleHeader() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderPropertyList
	// Experimental.
	SingleHeaderInput() interface{}
	// Experimental.
	SingleQueryArgument() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentPropertyList
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
	UriFragment() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchUriFragmentPropertyList
	// Experimental.
	UriFragmentInput() interface{}
	// Experimental.
	UriPath() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchUriPathPropertyList
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

// The jsii proxy struct for TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference
type jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) AllQueryArguments() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsPropertyList {
	var returns TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsPropertyList
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) AllQueryArgumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Body() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchBodyPropertyList {
	var returns TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchBodyPropertyList
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) BodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Cookies() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchCookiesPropertyList {
	var returns TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchCookiesPropertyList
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) CookiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) HeaderOrder() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchHeaderOrderPropertyList {
	var returns TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchHeaderOrderPropertyList
	_jsii_.Get(
		j,
		"headerOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) HeaderOrderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Headers() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchHeadersPropertyList {
	var returns TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchHeadersPropertyList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Ja3Fingerprint() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJa3FingerprintPropertyList {
	var returns TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJa3FingerprintPropertyList
	_jsii_.Get(
		j,
		"ja3Fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Ja3FingerprintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ja3FingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Ja4Fingerprint() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJa4FingerprintPropertyList {
	var returns TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJa4FingerprintPropertyList
	_jsii_.Get(
		j,
		"ja4Fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Ja4FingerprintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ja4FingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) JsonBody() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyPropertyList {
	var returns TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyPropertyList
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) JsonBodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Method() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList {
	var returns TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) MethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) QueryString() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchQueryStringPropertyList {
	var returns TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchQueryStringPropertyList
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) SingleHeader() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderPropertyList {
	var returns TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderPropertyList
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) SingleHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) SingleQueryArgument() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentPropertyList {
	var returns TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentPropertyList
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) SingleQueryArgumentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) UriFragment() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchUriFragmentPropertyList {
	var returns TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchUriFragmentPropertyList
	_jsii_.Get(
		j,
		"uriFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) UriFragmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uriFragmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) UriPath() TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchUriPathPropertyList {
	var returns TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchUriPathPropertyList
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) UriPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference_Override(t TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutAllQueryArguments(value interface{}) {
	if err := t.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutBody(value interface{}) {
	if err := t.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBody",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutCookies(value interface{}) {
	if err := t.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCookies",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutHeaderOrder(value interface{}) {
	if err := t.validatePutHeaderOrderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHeaderOrder",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutHeaders(value interface{}) {
	if err := t.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHeaders",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutJa3Fingerprint(value interface{}) {
	if err := t.validatePutJa3FingerprintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJa3Fingerprint",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutJa4Fingerprint(value interface{}) {
	if err := t.validatePutJa4FingerprintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJa4Fingerprint",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutJsonBody(value interface{}) {
	if err := t.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutMethod(value interface{}) {
	if err := t.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMethod",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutQueryString(value interface{}) {
	if err := t.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putQueryString",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutSingleHeader(value interface{}) {
	if err := t.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutSingleQueryArgument(value interface{}) {
	if err := t.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutUriFragment(value interface{}) {
	if err := t.validatePutUriFragmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUriFragment",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutUriPath(value interface{}) {
	if err := t.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUriPath",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		t,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		t,
		"resetBody",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		t,
		"resetCookies",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetHeaderOrder() {
	_jsii_.InvokeVoid(
		t,
		"resetHeaderOrder",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		t,
		"resetHeaders",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetJa3Fingerprint() {
	_jsii_.InvokeVoid(
		t,
		"resetJa3Fingerprint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetJa4Fingerprint() {
	_jsii_.InvokeVoid(
		t,
		"resetJa4Fingerprint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		t,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		t,
		"resetMethod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		t,
		"resetQueryString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		t,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetUriFragment() {
	_jsii_.InvokeVoid(
		t,
		"resetUriFragment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		t,
		"resetUriPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

