package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_CustomKeysPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Asn() TfWebAclRule_AsnPropertyList
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
	Cookie() TfWebAclRule_CookiePropertyList
	// Experimental.
	CookieInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	ForwardedIp() TfWebAclRule_ForwardedIpPropertyList
	// Experimental.
	ForwardedIpInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	Header() TfWebAclRule_StatementRateBasedStatementCustomKeysHeaderPropertyList
	// Experimental.
	HeaderInput() interface{}
	// Experimental.
	HttpMethod() TfWebAclRule_HttpMethodPropertyList
	// Experimental.
	HttpMethodInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Ip() TfWebAclRule_IpPropertyList
	// Experimental.
	IpInput() interface{}
	// Experimental.
	Ja3Fingerprint() TfWebAclRule_StatementRateBasedStatementCustomKeysJa3FingerprintPropertyList
	// Experimental.
	Ja3FingerprintInput() interface{}
	// Experimental.
	Ja4Fingerprint() TfWebAclRule_StatementRateBasedStatementCustomKeysJa4FingerprintPropertyList
	// Experimental.
	Ja4FingerprintInput() interface{}
	// Experimental.
	LabelNamespace() TfWebAclRule_LabelNamespacePropertyList
	// Experimental.
	LabelNamespaceInput() interface{}
	// Experimental.
	QueryArgument() TfWebAclRule_QueryArgumentPropertyList
	// Experimental.
	QueryArgumentInput() interface{}
	// Experimental.
	QueryString() TfWebAclRule_StatementRateBasedStatementCustomKeysQueryStringPropertyList
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
	UriPath() TfWebAclRule_StatementRateBasedStatementCustomKeysUriPathPropertyList
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

// The jsii proxy struct for TfWebAclRule_CustomKeysPropertyOutputReference
type jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) Asn() TfWebAclRule_AsnPropertyList {
	var returns TfWebAclRule_AsnPropertyList
	_jsii_.Get(
		j,
		"asn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) AsnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"asnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) Cookie() TfWebAclRule_CookiePropertyList {
	var returns TfWebAclRule_CookiePropertyList
	_jsii_.Get(
		j,
		"cookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) CookieInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ForwardedIp() TfWebAclRule_ForwardedIpPropertyList {
	var returns TfWebAclRule_ForwardedIpPropertyList
	_jsii_.Get(
		j,
		"forwardedIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ForwardedIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardedIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) Header() TfWebAclRule_StatementRateBasedStatementCustomKeysHeaderPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementCustomKeysHeaderPropertyList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) HttpMethod() TfWebAclRule_HttpMethodPropertyList {
	var returns TfWebAclRule_HttpMethodPropertyList
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) HttpMethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) Ip() TfWebAclRule_IpPropertyList {
	var returns TfWebAclRule_IpPropertyList
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) IpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) Ja3Fingerprint() TfWebAclRule_StatementRateBasedStatementCustomKeysJa3FingerprintPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementCustomKeysJa3FingerprintPropertyList
	_jsii_.Get(
		j,
		"ja3Fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) Ja3FingerprintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ja3FingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) Ja4Fingerprint() TfWebAclRule_StatementRateBasedStatementCustomKeysJa4FingerprintPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementCustomKeysJa4FingerprintPropertyList
	_jsii_.Get(
		j,
		"ja4Fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) Ja4FingerprintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ja4FingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) LabelNamespace() TfWebAclRule_LabelNamespacePropertyList {
	var returns TfWebAclRule_LabelNamespacePropertyList
	_jsii_.Get(
		j,
		"labelNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) LabelNamespaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) QueryArgument() TfWebAclRule_QueryArgumentPropertyList {
	var returns TfWebAclRule_QueryArgumentPropertyList
	_jsii_.Get(
		j,
		"queryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) QueryArgumentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) QueryString() TfWebAclRule_StatementRateBasedStatementCustomKeysQueryStringPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementCustomKeysQueryStringPropertyList
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) UriPath() TfWebAclRule_StatementRateBasedStatementCustomKeysUriPathPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementCustomKeysUriPathPropertyList
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) UriPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_CustomKeysPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRule_CustomKeysPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_CustomKeysPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.CustomKeysPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_CustomKeysPropertyOutputReference_Override(t TfWebAclRule_CustomKeysPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.CustomKeysPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) PutAsn(value interface{}) {
	if err := t.validatePutAsnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAsn",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) PutCookie(value interface{}) {
	if err := t.validatePutCookieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCookie",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) PutForwardedIp(value interface{}) {
	if err := t.validatePutForwardedIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putForwardedIp",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) PutHeader(value interface{}) {
	if err := t.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) PutHttpMethod(value interface{}) {
	if err := t.validatePutHttpMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHttpMethod",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) PutIp(value interface{}) {
	if err := t.validatePutIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIp",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) PutJa3Fingerprint(value interface{}) {
	if err := t.validatePutJa3FingerprintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJa3Fingerprint",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) PutJa4Fingerprint(value interface{}) {
	if err := t.validatePutJa4FingerprintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJa4Fingerprint",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) PutLabelNamespace(value interface{}) {
	if err := t.validatePutLabelNamespaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLabelNamespace",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) PutQueryArgument(value interface{}) {
	if err := t.validatePutQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putQueryArgument",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) PutQueryString(value interface{}) {
	if err := t.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putQueryString",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) PutUriPath(value interface{}) {
	if err := t.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUriPath",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ResetAsn() {
	_jsii_.InvokeVoid(
		t,
		"resetAsn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ResetCookie() {
	_jsii_.InvokeVoid(
		t,
		"resetCookie",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ResetForwardedIp() {
	_jsii_.InvokeVoid(
		t,
		"resetForwardedIp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		t,
		"resetIp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ResetJa3Fingerprint() {
	_jsii_.InvokeVoid(
		t,
		"resetJa3Fingerprint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ResetJa4Fingerprint() {
	_jsii_.InvokeVoid(
		t,
		"resetJa4Fingerprint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ResetLabelNamespace() {
	_jsii_.InvokeVoid(
		t,
		"resetLabelNamespace",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ResetQueryArgument() {
	_jsii_.InvokeVoid(
		t,
		"resetQueryArgument",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		t,
		"resetQueryString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		t,
		"resetUriPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_CustomKeysPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

