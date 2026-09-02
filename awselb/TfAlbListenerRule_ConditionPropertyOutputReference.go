package awselb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAlbListenerRule_ConditionPropertyOutputReference interface {
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
	HostHeader() TfAlbListenerRule_HostHeaderPropertyOutputReference
	// Experimental.
	HostHeaderInput() *TfAlbListenerRule_HostHeaderProperty
	// Experimental.
	HttpHeader() TfAlbListenerRule_HttpHeaderPropertyOutputReference
	// Experimental.
	HttpHeaderInput() *TfAlbListenerRule_HttpHeaderProperty
	// Experimental.
	HttpRequestMethod() TfAlbListenerRule_HttpRequestMethodPropertyOutputReference
	// Experimental.
	HttpRequestMethodInput() *TfAlbListenerRule_HttpRequestMethodProperty
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PathPattern() TfAlbListenerRule_PathPatternPropertyOutputReference
	// Experimental.
	PathPatternInput() *TfAlbListenerRule_PathPatternProperty
	// Experimental.
	QueryString() TfAlbListenerRule_QueryStringPropertyList
	// Experimental.
	QueryStringInput() interface{}
	// Experimental.
	SourceIp() TfAlbListenerRule_SourceIpPropertyOutputReference
	// Experimental.
	SourceIpInput() *TfAlbListenerRule_SourceIpProperty
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
	PutHostHeader(value *TfAlbListenerRule_HostHeaderProperty)
	// Experimental.
	PutHttpHeader(value *TfAlbListenerRule_HttpHeaderProperty)
	// Experimental.
	PutHttpRequestMethod(value *TfAlbListenerRule_HttpRequestMethodProperty)
	// Experimental.
	PutPathPattern(value *TfAlbListenerRule_PathPatternProperty)
	// Experimental.
	PutQueryString(value interface{})
	// Experimental.
	PutSourceIp(value *TfAlbListenerRule_SourceIpProperty)
	// Experimental.
	ResetHostHeader()
	// Experimental.
	ResetHttpHeader()
	// Experimental.
	ResetHttpRequestMethod()
	// Experimental.
	ResetPathPattern()
	// Experimental.
	ResetQueryString()
	// Experimental.
	ResetSourceIp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfAlbListenerRule_ConditionPropertyOutputReference
type jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) HostHeader() TfAlbListenerRule_HostHeaderPropertyOutputReference {
	var returns TfAlbListenerRule_HostHeaderPropertyOutputReference
	_jsii_.Get(
		j,
		"hostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) HostHeaderInput() *TfAlbListenerRule_HostHeaderProperty {
	var returns *TfAlbListenerRule_HostHeaderProperty
	_jsii_.Get(
		j,
		"hostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) HttpHeader() TfAlbListenerRule_HttpHeaderPropertyOutputReference {
	var returns TfAlbListenerRule_HttpHeaderPropertyOutputReference
	_jsii_.Get(
		j,
		"httpHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) HttpHeaderInput() *TfAlbListenerRule_HttpHeaderProperty {
	var returns *TfAlbListenerRule_HttpHeaderProperty
	_jsii_.Get(
		j,
		"httpHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) HttpRequestMethod() TfAlbListenerRule_HttpRequestMethodPropertyOutputReference {
	var returns TfAlbListenerRule_HttpRequestMethodPropertyOutputReference
	_jsii_.Get(
		j,
		"httpRequestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) HttpRequestMethodInput() *TfAlbListenerRule_HttpRequestMethodProperty {
	var returns *TfAlbListenerRule_HttpRequestMethodProperty
	_jsii_.Get(
		j,
		"httpRequestMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) PathPattern() TfAlbListenerRule_PathPatternPropertyOutputReference {
	var returns TfAlbListenerRule_PathPatternPropertyOutputReference
	_jsii_.Get(
		j,
		"pathPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) PathPatternInput() *TfAlbListenerRule_PathPatternProperty {
	var returns *TfAlbListenerRule_PathPatternProperty
	_jsii_.Get(
		j,
		"pathPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) QueryString() TfAlbListenerRule_QueryStringPropertyList {
	var returns TfAlbListenerRule_QueryStringPropertyList
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) SourceIp() TfAlbListenerRule_SourceIpPropertyOutputReference {
	var returns TfAlbListenerRule_SourceIpPropertyOutputReference
	_jsii_.Get(
		j,
		"sourceIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) SourceIpInput() *TfAlbListenerRule_SourceIpProperty {
	var returns *TfAlbListenerRule_SourceIpProperty
	_jsii_.Get(
		j,
		"sourceIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAlbListenerRule_ConditionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfAlbListenerRule_ConditionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAlbListenerRule_ConditionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elb.TfAlbListenerRule.ConditionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAlbListenerRule_ConditionPropertyOutputReference_Override(t TfAlbListenerRule_ConditionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.TfAlbListenerRule.ConditionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) PutHostHeader(value *TfAlbListenerRule_HostHeaderProperty) {
	if err := t.validatePutHostHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHostHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) PutHttpHeader(value *TfAlbListenerRule_HttpHeaderProperty) {
	if err := t.validatePutHttpHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHttpHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) PutHttpRequestMethod(value *TfAlbListenerRule_HttpRequestMethodProperty) {
	if err := t.validatePutHttpRequestMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHttpRequestMethod",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) PutPathPattern(value *TfAlbListenerRule_PathPatternProperty) {
	if err := t.validatePutPathPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPathPattern",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) PutQueryString(value interface{}) {
	if err := t.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putQueryString",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) PutSourceIp(value *TfAlbListenerRule_SourceIpProperty) {
	if err := t.validatePutSourceIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSourceIp",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) ResetHostHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetHostHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) ResetHttpHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) ResetHttpRequestMethod() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpRequestMethod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) ResetPathPattern() {
	_jsii_.InvokeVoid(
		t,
		"resetPathPattern",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		t,
		"resetQueryString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) ResetSourceIp() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceIp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAlbListenerRule_ConditionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

