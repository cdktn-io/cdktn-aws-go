package waf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/waf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/waf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllQueryArguments() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchAllQueryArgumentsPropertyList
	// Experimental.
	AllQueryArgumentsInput() interface{}
	// Experimental.
	Body() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchBodyPropertyList
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
	Cookies() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchCookiesPropertyList
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
	HeaderOrder() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchHeaderOrderPropertyList
	// Experimental.
	HeaderOrderInput() interface{}
	// Experimental.
	Headers() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchHeadersPropertyList
	// Experimental.
	HeadersInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Ja3Fingerprint() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchJa3FingerprintPropertyList
	// Experimental.
	Ja3FingerprintInput() interface{}
	// Experimental.
	Ja4Fingerprint() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchJa4FingerprintPropertyList
	// Experimental.
	Ja4FingerprintInput() interface{}
	// Experimental.
	JsonBody() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchJsonBodyPropertyList
	// Experimental.
	JsonBodyInput() interface{}
	// Experimental.
	Method() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchMethodPropertyList
	// Experimental.
	MethodInput() interface{}
	// Experimental.
	QueryString() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchQueryStringPropertyList
	// Experimental.
	QueryStringInput() interface{}
	// Experimental.
	SingleHeader() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchSingleHeaderPropertyList
	// Experimental.
	SingleHeaderInput() interface{}
	// Experimental.
	SingleQueryArgument() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchSingleQueryArgumentPropertyList
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
	UriFragment() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchUriFragmentPropertyList
	// Experimental.
	UriFragmentInput() interface{}
	// Experimental.
	UriPath() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchUriPathPropertyList
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

// The jsii proxy struct for AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference
type jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) AllQueryArguments() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchAllQueryArgumentsPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchAllQueryArgumentsPropertyList
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) AllQueryArgumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) Body() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchBodyPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchBodyPropertyList
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) BodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) Cookies() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchCookiesPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchCookiesPropertyList
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) CookiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) HeaderOrder() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchHeaderOrderPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchHeaderOrderPropertyList
	_jsii_.Get(
		j,
		"headerOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) HeaderOrderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) Headers() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchHeadersPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchHeadersPropertyList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) Ja3Fingerprint() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchJa3FingerprintPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchJa3FingerprintPropertyList
	_jsii_.Get(
		j,
		"ja3Fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) Ja3FingerprintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ja3FingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) Ja4Fingerprint() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchJa4FingerprintPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchJa4FingerprintPropertyList
	_jsii_.Get(
		j,
		"ja4Fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) Ja4FingerprintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ja4FingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) JsonBody() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchJsonBodyPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchJsonBodyPropertyList
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) JsonBodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) Method() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchMethodPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchMethodPropertyList
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) MethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) QueryString() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchQueryStringPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchQueryStringPropertyList
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) SingleHeader() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchSingleHeaderPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchSingleHeaderPropertyList
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) SingleHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) SingleQueryArgument() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchSingleQueryArgumentPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchSingleQueryArgumentPropertyList
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) SingleQueryArgumentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) UriFragment() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchUriFragmentPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchUriFragmentPropertyList
	_jsii_.Get(
		j,
		"uriFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) UriFragmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uriFragmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) UriPath() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchUriPathPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchUriPathPropertyList
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) UriPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRule.StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference_Override(a AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRule.StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) PutAllQueryArguments(value interface{}) {
	if err := a.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) PutBody(value interface{}) {
	if err := a.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBody",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) PutCookies(value interface{}) {
	if err := a.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCookies",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) PutHeaderOrder(value interface{}) {
	if err := a.validatePutHeaderOrderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHeaderOrder",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) PutHeaders(value interface{}) {
	if err := a.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHeaders",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) PutJa3Fingerprint(value interface{}) {
	if err := a.validatePutJa3FingerprintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJa3Fingerprint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) PutJa4Fingerprint(value interface{}) {
	if err := a.validatePutJa4FingerprintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJa4Fingerprint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) PutJsonBody(value interface{}) {
	if err := a.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) PutMethod(value interface{}) {
	if err := a.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMethod",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) PutQueryString(value interface{}) {
	if err := a.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueryString",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) PutSingleHeader(value interface{}) {
	if err := a.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) PutSingleQueryArgument(value interface{}) {
	if err := a.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) PutUriFragment(value interface{}) {
	if err := a.validatePutUriFragmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUriFragment",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) PutUriPath(value interface{}) {
	if err := a.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUriPath",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		a,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		a,
		"resetBody",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		a,
		"resetCookies",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ResetHeaderOrder() {
	_jsii_.InvokeVoid(
		a,
		"resetHeaderOrder",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ResetJa3Fingerprint() {
	_jsii_.InvokeVoid(
		a,
		"resetJa3Fingerprint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ResetJa4Fingerprint() {
	_jsii_.InvokeVoid(
		a,
		"resetJa4Fingerprint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		a,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		a,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ResetUriFragment() {
	_jsii_.InvokeVoid(
		a,
		"resetUriFragment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		a,
		"resetUriPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

