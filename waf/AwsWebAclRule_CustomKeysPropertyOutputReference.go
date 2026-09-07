package waf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/waf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/waf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWebAclRule_CustomKeysPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Asn() AwsWebAclRule_AsnPropertyList
	// Experimental.
	AsnInput() interface{}
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
	Cookie() AwsWebAclRule_CookiePropertyList
	// Experimental.
	CookieInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	ForwardedIp() AwsWebAclRule_ForwardedIpPropertyList
	// Experimental.
	ForwardedIpInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	Header() AwsWebAclRule_StatementRateBasedStatementCustomKeysHeaderPropertyList
	// Experimental.
	HeaderInput() interface{}
	// Experimental.
	HttpMethod() AwsWebAclRule_HttpMethodPropertyList
	// Experimental.
	HttpMethodInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Ip() AwsWebAclRule_IpPropertyList
	// Experimental.
	IpInput() interface{}
	// Experimental.
	Ja3Fingerprint() AwsWebAclRule_StatementRateBasedStatementCustomKeysJa3FingerprintPropertyList
	// Experimental.
	Ja3FingerprintInput() interface{}
	// Experimental.
	Ja4Fingerprint() AwsWebAclRule_StatementRateBasedStatementCustomKeysJa4FingerprintPropertyList
	// Experimental.
	Ja4FingerprintInput() interface{}
	// Experimental.
	LabelNamespace() AwsWebAclRule_LabelNamespacePropertyList
	// Experimental.
	LabelNamespaceInput() interface{}
	// Experimental.
	QueryArgument() AwsWebAclRule_QueryArgumentPropertyList
	// Experimental.
	QueryArgumentInput() interface{}
	// Experimental.
	QueryString() AwsWebAclRule_StatementRateBasedStatementCustomKeysQueryStringPropertyList
	// Experimental.
	QueryStringInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UriPath() AwsWebAclRule_StatementRateBasedStatementCustomKeysUriPathPropertyList
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
	PutAsn(value interface{})
	// Experimental.
	PutCookie(value interface{})
	// Experimental.
	PutForwardedIp(value interface{})
	// Experimental.
	PutHeader(value interface{})
	// Experimental.
	PutHttpMethod(value interface{})
	// Experimental.
	PutIp(value interface{})
	// Experimental.
	PutJa3Fingerprint(value interface{})
	// Experimental.
	PutJa4Fingerprint(value interface{})
	// Experimental.
	PutLabelNamespace(value interface{})
	// Experimental.
	PutQueryArgument(value interface{})
	// Experimental.
	PutQueryString(value interface{})
	// Experimental.
	PutUriPath(value interface{})
	// Experimental.
	ResetAsn()
	// Experimental.
	ResetCookie()
	// Experimental.
	ResetForwardedIp()
	// Experimental.
	ResetHeader()
	// Experimental.
	ResetHttpMethod()
	// Experimental.
	ResetIp()
	// Experimental.
	ResetJa3Fingerprint()
	// Experimental.
	ResetJa4Fingerprint()
	// Experimental.
	ResetLabelNamespace()
	// Experimental.
	ResetQueryArgument()
	// Experimental.
	ResetQueryString()
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

// The jsii proxy struct for AwsWebAclRule_CustomKeysPropertyOutputReference
type jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) Asn() AwsWebAclRule_AsnPropertyList {
	var returns AwsWebAclRule_AsnPropertyList
	_jsii_.Get(
		j,
		"asn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) AsnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"asnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) Cookie() AwsWebAclRule_CookiePropertyList {
	var returns AwsWebAclRule_CookiePropertyList
	_jsii_.Get(
		j,
		"cookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) CookieInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ForwardedIp() AwsWebAclRule_ForwardedIpPropertyList {
	var returns AwsWebAclRule_ForwardedIpPropertyList
	_jsii_.Get(
		j,
		"forwardedIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ForwardedIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardedIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) Header() AwsWebAclRule_StatementRateBasedStatementCustomKeysHeaderPropertyList {
	var returns AwsWebAclRule_StatementRateBasedStatementCustomKeysHeaderPropertyList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) HttpMethod() AwsWebAclRule_HttpMethodPropertyList {
	var returns AwsWebAclRule_HttpMethodPropertyList
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) HttpMethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) Ip() AwsWebAclRule_IpPropertyList {
	var returns AwsWebAclRule_IpPropertyList
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) IpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) Ja3Fingerprint() AwsWebAclRule_StatementRateBasedStatementCustomKeysJa3FingerprintPropertyList {
	var returns AwsWebAclRule_StatementRateBasedStatementCustomKeysJa3FingerprintPropertyList
	_jsii_.Get(
		j,
		"ja3Fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) Ja3FingerprintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ja3FingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) Ja4Fingerprint() AwsWebAclRule_StatementRateBasedStatementCustomKeysJa4FingerprintPropertyList {
	var returns AwsWebAclRule_StatementRateBasedStatementCustomKeysJa4FingerprintPropertyList
	_jsii_.Get(
		j,
		"ja4Fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) Ja4FingerprintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ja4FingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) LabelNamespace() AwsWebAclRule_LabelNamespacePropertyList {
	var returns AwsWebAclRule_LabelNamespacePropertyList
	_jsii_.Get(
		j,
		"labelNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) LabelNamespaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) QueryArgument() AwsWebAclRule_QueryArgumentPropertyList {
	var returns AwsWebAclRule_QueryArgumentPropertyList
	_jsii_.Get(
		j,
		"queryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) QueryArgumentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) QueryString() AwsWebAclRule_StatementRateBasedStatementCustomKeysQueryStringPropertyList {
	var returns AwsWebAclRule_StatementRateBasedStatementCustomKeysQueryStringPropertyList
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) UriPath() AwsWebAclRule_StatementRateBasedStatementCustomKeysUriPathPropertyList {
	var returns AwsWebAclRule_StatementRateBasedStatementCustomKeysUriPathPropertyList
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) UriPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWebAclRule_CustomKeysPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsWebAclRule_CustomKeysPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWebAclRule_CustomKeysPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRule.CustomKeysPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWebAclRule_CustomKeysPropertyOutputReference_Override(a AwsWebAclRule_CustomKeysPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRule.CustomKeysPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) PutAsn(value interface{}) {
	if err := a.validatePutAsnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAsn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) PutCookie(value interface{}) {
	if err := a.validatePutCookieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCookie",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) PutForwardedIp(value interface{}) {
	if err := a.validatePutForwardedIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putForwardedIp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) PutHeader(value interface{}) {
	if err := a.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHeader",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) PutHttpMethod(value interface{}) {
	if err := a.validatePutHttpMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpMethod",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) PutIp(value interface{}) {
	if err := a.validatePutIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) PutJa3Fingerprint(value interface{}) {
	if err := a.validatePutJa3FingerprintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJa3Fingerprint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) PutJa4Fingerprint(value interface{}) {
	if err := a.validatePutJa4FingerprintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJa4Fingerprint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) PutLabelNamespace(value interface{}) {
	if err := a.validatePutLabelNamespaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLabelNamespace",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) PutQueryArgument(value interface{}) {
	if err := a.validatePutQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueryArgument",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) PutQueryString(value interface{}) {
	if err := a.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueryString",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) PutUriPath(value interface{}) {
	if err := a.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUriPath",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ResetAsn() {
	_jsii_.InvokeVoid(
		a,
		"resetAsn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ResetCookie() {
	_jsii_.InvokeVoid(
		a,
		"resetCookie",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ResetForwardedIp() {
	_jsii_.InvokeVoid(
		a,
		"resetForwardedIp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		a,
		"resetIp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ResetJa3Fingerprint() {
	_jsii_.InvokeVoid(
		a,
		"resetJa3Fingerprint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ResetJa4Fingerprint() {
	_jsii_.InvokeVoid(
		a,
		"resetJa4Fingerprint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ResetLabelNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetLabelNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ResetQueryArgument() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryArgument",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		a,
		"resetUriPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWebAclRule_CustomKeysPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

