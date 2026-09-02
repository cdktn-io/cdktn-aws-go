package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllQueryArguments() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsPropertyList
	// Experimental.
	AllQueryArgumentsInput() interface{}
	// Experimental.
	Body() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchBodyPropertyList
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
	Cookies() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchCookiesPropertyList
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
	HeaderOrder() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeaderOrderPropertyList
	// Experimental.
	HeaderOrderInput() interface{}
	// Experimental.
	Headers() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersPropertyList
	// Experimental.
	HeadersInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Ja3Fingerprint() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJa3FingerprintPropertyList
	// Experimental.
	Ja3FingerprintInput() interface{}
	// Experimental.
	Ja4Fingerprint() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJa4FingerprintPropertyList
	// Experimental.
	Ja4FingerprintInput() interface{}
	// Experimental.
	JsonBody() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyPropertyList
	// Experimental.
	JsonBodyInput() interface{}
	// Experimental.
	Method() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList
	// Experimental.
	MethodInput() interface{}
	// Experimental.
	QueryString() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringPropertyList
	// Experimental.
	QueryStringInput() interface{}
	// Experimental.
	SingleHeader() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderPropertyList
	// Experimental.
	SingleHeaderInput() interface{}
	// Experimental.
	SingleQueryArgument() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentPropertyList
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
	UriFragment() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriFragmentPropertyList
	// Experimental.
	UriFragmentInput() interface{}
	// Experimental.
	UriPath() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriPathPropertyList
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

// The jsii proxy struct for TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference
type jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) AllQueryArguments() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsPropertyList
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) AllQueryArgumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Body() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchBodyPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchBodyPropertyList
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) BodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Cookies() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchCookiesPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchCookiesPropertyList
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) CookiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) HeaderOrder() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeaderOrderPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeaderOrderPropertyList
	_jsii_.Get(
		j,
		"headerOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) HeaderOrderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Headers() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersPropertyList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Ja3Fingerprint() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJa3FingerprintPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJa3FingerprintPropertyList
	_jsii_.Get(
		j,
		"ja3Fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Ja3FingerprintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ja3FingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Ja4Fingerprint() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJa4FingerprintPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJa4FingerprintPropertyList
	_jsii_.Get(
		j,
		"ja4Fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Ja4FingerprintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ja4FingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) JsonBody() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyPropertyList
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) JsonBodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Method() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) MethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) QueryString() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringPropertyList
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) SingleHeader() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderPropertyList
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) SingleHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) SingleQueryArgument() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentPropertyList
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) SingleQueryArgumentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) UriFragment() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriFragmentPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriFragmentPropertyList
	_jsii_.Get(
		j,
		"uriFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) UriFragmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uriFragmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) UriPath() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriPathPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriPathPropertyList
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) UriPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference_Override(t TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutAllQueryArguments(value interface{}) {
	if err := t.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutBody(value interface{}) {
	if err := t.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBody",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutCookies(value interface{}) {
	if err := t.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCookies",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutHeaderOrder(value interface{}) {
	if err := t.validatePutHeaderOrderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHeaderOrder",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutHeaders(value interface{}) {
	if err := t.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHeaders",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutJa3Fingerprint(value interface{}) {
	if err := t.validatePutJa3FingerprintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJa3Fingerprint",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutJa4Fingerprint(value interface{}) {
	if err := t.validatePutJa4FingerprintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJa4Fingerprint",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutJsonBody(value interface{}) {
	if err := t.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutMethod(value interface{}) {
	if err := t.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMethod",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutQueryString(value interface{}) {
	if err := t.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putQueryString",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutSingleHeader(value interface{}) {
	if err := t.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutSingleQueryArgument(value interface{}) {
	if err := t.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutUriFragment(value interface{}) {
	if err := t.validatePutUriFragmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUriFragment",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) PutUriPath(value interface{}) {
	if err := t.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUriPath",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		t,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		t,
		"resetBody",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		t,
		"resetCookies",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetHeaderOrder() {
	_jsii_.InvokeVoid(
		t,
		"resetHeaderOrder",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		t,
		"resetHeaders",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetJa3Fingerprint() {
	_jsii_.InvokeVoid(
		t,
		"resetJa3Fingerprint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetJa4Fingerprint() {
	_jsii_.InvokeVoid(
		t,
		"resetJa4Fingerprint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		t,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		t,
		"resetMethod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		t,
		"resetQueryString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		t,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetUriFragment() {
	_jsii_.InvokeVoid(
		t,
		"resetUriFragment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		t,
		"resetUriPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

