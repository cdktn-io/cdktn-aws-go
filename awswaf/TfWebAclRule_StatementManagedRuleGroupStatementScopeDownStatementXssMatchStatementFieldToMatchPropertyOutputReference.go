package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllQueryArguments() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchAllQueryArgumentsPropertyList
	// Experimental.
	AllQueryArgumentsInput() interface{}
	// Experimental.
	Body() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchBodyPropertyList
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
	Cookies() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchCookiesPropertyList
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
	HeaderOrder() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchHeaderOrderPropertyList
	// Experimental.
	HeaderOrderInput() interface{}
	// Experimental.
	Headers() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchHeadersPropertyList
	// Experimental.
	HeadersInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Ja3Fingerprint() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchJa3FingerprintPropertyList
	// Experimental.
	Ja3FingerprintInput() interface{}
	// Experimental.
	Ja4Fingerprint() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchJa4FingerprintPropertyList
	// Experimental.
	Ja4FingerprintInput() interface{}
	// Experimental.
	JsonBody() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchJsonBodyPropertyList
	// Experimental.
	JsonBodyInput() interface{}
	// Experimental.
	Method() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchMethodPropertyList
	// Experimental.
	MethodInput() interface{}
	// Experimental.
	QueryString() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchQueryStringPropertyList
	// Experimental.
	QueryStringInput() interface{}
	// Experimental.
	SingleHeader() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchSingleHeaderPropertyList
	// Experimental.
	SingleHeaderInput() interface{}
	// Experimental.
	SingleQueryArgument() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchSingleQueryArgumentPropertyList
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
	UriFragment() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchUriFragmentPropertyList
	// Experimental.
	UriFragmentInput() interface{}
	// Experimental.
	UriPath() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchUriPathPropertyList
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

// The jsii proxy struct for TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference
type jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) AllQueryArguments() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchAllQueryArgumentsPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchAllQueryArgumentsPropertyList
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) AllQueryArgumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) Body() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchBodyPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchBodyPropertyList
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) BodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) Cookies() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchCookiesPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchCookiesPropertyList
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) CookiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) HeaderOrder() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchHeaderOrderPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchHeaderOrderPropertyList
	_jsii_.Get(
		j,
		"headerOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) HeaderOrderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) Headers() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchHeadersPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchHeadersPropertyList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) Ja3Fingerprint() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchJa3FingerprintPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchJa3FingerprintPropertyList
	_jsii_.Get(
		j,
		"ja3Fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) Ja3FingerprintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ja3FingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) Ja4Fingerprint() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchJa4FingerprintPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchJa4FingerprintPropertyList
	_jsii_.Get(
		j,
		"ja4Fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) Ja4FingerprintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ja4FingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) JsonBody() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchJsonBodyPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchJsonBodyPropertyList
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) JsonBodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) Method() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchMethodPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchMethodPropertyList
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) MethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) QueryString() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchQueryStringPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchQueryStringPropertyList
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) SingleHeader() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchSingleHeaderPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchSingleHeaderPropertyList
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) SingleHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) SingleQueryArgument() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchSingleQueryArgumentPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchSingleQueryArgumentPropertyList
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) SingleQueryArgumentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) UriFragment() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchUriFragmentPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchUriFragmentPropertyList
	_jsii_.Get(
		j,
		"uriFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) UriFragmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uriFragmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) UriPath() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchUriPathPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchUriPathPropertyList
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) UriPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference_Override(t TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) PutAllQueryArguments(value interface{}) {
	if err := t.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) PutBody(value interface{}) {
	if err := t.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBody",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) PutCookies(value interface{}) {
	if err := t.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCookies",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) PutHeaderOrder(value interface{}) {
	if err := t.validatePutHeaderOrderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHeaderOrder",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) PutHeaders(value interface{}) {
	if err := t.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHeaders",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) PutJa3Fingerprint(value interface{}) {
	if err := t.validatePutJa3FingerprintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJa3Fingerprint",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) PutJa4Fingerprint(value interface{}) {
	if err := t.validatePutJa4FingerprintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJa4Fingerprint",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) PutJsonBody(value interface{}) {
	if err := t.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) PutMethod(value interface{}) {
	if err := t.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMethod",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) PutQueryString(value interface{}) {
	if err := t.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putQueryString",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) PutSingleHeader(value interface{}) {
	if err := t.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) PutSingleQueryArgument(value interface{}) {
	if err := t.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) PutUriFragment(value interface{}) {
	if err := t.validatePutUriFragmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUriFragment",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) PutUriPath(value interface{}) {
	if err := t.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUriPath",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		t,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		t,
		"resetBody",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		t,
		"resetCookies",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetHeaderOrder() {
	_jsii_.InvokeVoid(
		t,
		"resetHeaderOrder",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		t,
		"resetHeaders",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetJa3Fingerprint() {
	_jsii_.InvokeVoid(
		t,
		"resetJa3Fingerprint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetJa4Fingerprint() {
	_jsii_.InvokeVoid(
		t,
		"resetJa4Fingerprint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		t,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		t,
		"resetMethod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		t,
		"resetQueryString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		t,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetUriFragment() {
	_jsii_.InvokeVoid(
		t,
		"resetUriFragment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		t,
		"resetUriPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

