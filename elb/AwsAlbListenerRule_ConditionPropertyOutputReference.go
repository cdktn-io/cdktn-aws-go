package elb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAlbListenerRule_ConditionPropertyOutputReference interface {
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
	HostHeader() AwsAlbListenerRule_HostHeaderPropertyOutputReference
	// Experimental.
	HostHeaderInput() *AwsAlbListenerRule_HostHeaderProperty
	// Experimental.
	HttpHeader() AwsAlbListenerRule_HttpHeaderPropertyOutputReference
	// Experimental.
	HttpHeaderInput() *AwsAlbListenerRule_HttpHeaderProperty
	// Experimental.
	HttpRequestMethod() AwsAlbListenerRule_HttpRequestMethodPropertyOutputReference
	// Experimental.
	HttpRequestMethodInput() *AwsAlbListenerRule_HttpRequestMethodProperty
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PathPattern() AwsAlbListenerRule_PathPatternPropertyOutputReference
	// Experimental.
	PathPatternInput() *AwsAlbListenerRule_PathPatternProperty
	// Experimental.
	QueryString() AwsAlbListenerRule_QueryStringPropertyList
	// Experimental.
	QueryStringInput() interface{}
	// Experimental.
	SourceIp() AwsAlbListenerRule_SourceIpPropertyOutputReference
	// Experimental.
	SourceIpInput() *AwsAlbListenerRule_SourceIpProperty
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
	PutHostHeader(value *AwsAlbListenerRule_HostHeaderProperty)
	// Experimental.
	PutHttpHeader(value *AwsAlbListenerRule_HttpHeaderProperty)
	// Experimental.
	PutHttpRequestMethod(value *AwsAlbListenerRule_HttpRequestMethodProperty)
	// Experimental.
	PutPathPattern(value *AwsAlbListenerRule_PathPatternProperty)
	// Experimental.
	PutQueryString(value interface{})
	// Experimental.
	PutSourceIp(value *AwsAlbListenerRule_SourceIpProperty)
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

// The jsii proxy struct for AwsAlbListenerRule_ConditionPropertyOutputReference
type jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) HostHeader() AwsAlbListenerRule_HostHeaderPropertyOutputReference {
	var returns AwsAlbListenerRule_HostHeaderPropertyOutputReference
	_jsii_.Get(
		j,
		"hostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) HostHeaderInput() *AwsAlbListenerRule_HostHeaderProperty {
	var returns *AwsAlbListenerRule_HostHeaderProperty
	_jsii_.Get(
		j,
		"hostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) HttpHeader() AwsAlbListenerRule_HttpHeaderPropertyOutputReference {
	var returns AwsAlbListenerRule_HttpHeaderPropertyOutputReference
	_jsii_.Get(
		j,
		"httpHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) HttpHeaderInput() *AwsAlbListenerRule_HttpHeaderProperty {
	var returns *AwsAlbListenerRule_HttpHeaderProperty
	_jsii_.Get(
		j,
		"httpHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) HttpRequestMethod() AwsAlbListenerRule_HttpRequestMethodPropertyOutputReference {
	var returns AwsAlbListenerRule_HttpRequestMethodPropertyOutputReference
	_jsii_.Get(
		j,
		"httpRequestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) HttpRequestMethodInput() *AwsAlbListenerRule_HttpRequestMethodProperty {
	var returns *AwsAlbListenerRule_HttpRequestMethodProperty
	_jsii_.Get(
		j,
		"httpRequestMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) PathPattern() AwsAlbListenerRule_PathPatternPropertyOutputReference {
	var returns AwsAlbListenerRule_PathPatternPropertyOutputReference
	_jsii_.Get(
		j,
		"pathPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) PathPatternInput() *AwsAlbListenerRule_PathPatternProperty {
	var returns *AwsAlbListenerRule_PathPatternProperty
	_jsii_.Get(
		j,
		"pathPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) QueryString() AwsAlbListenerRule_QueryStringPropertyList {
	var returns AwsAlbListenerRule_QueryStringPropertyList
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) SourceIp() AwsAlbListenerRule_SourceIpPropertyOutputReference {
	var returns AwsAlbListenerRule_SourceIpPropertyOutputReference
	_jsii_.Get(
		j,
		"sourceIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) SourceIpInput() *AwsAlbListenerRule_SourceIpProperty {
	var returns *AwsAlbListenerRule_SourceIpProperty
	_jsii_.Get(
		j,
		"sourceIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAlbListenerRule_ConditionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsAlbListenerRule_ConditionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAlbListenerRule_ConditionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elb.AwsAlbListenerRule.ConditionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAlbListenerRule_ConditionPropertyOutputReference_Override(a AwsAlbListenerRule_ConditionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.AwsAlbListenerRule.ConditionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) PutHostHeader(value *AwsAlbListenerRule_HostHeaderProperty) {
	if err := a.validatePutHostHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHostHeader",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) PutHttpHeader(value *AwsAlbListenerRule_HttpHeaderProperty) {
	if err := a.validatePutHttpHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpHeader",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) PutHttpRequestMethod(value *AwsAlbListenerRule_HttpRequestMethodProperty) {
	if err := a.validatePutHttpRequestMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpRequestMethod",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) PutPathPattern(value *AwsAlbListenerRule_PathPatternProperty) {
	if err := a.validatePutPathPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPathPattern",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) PutQueryString(value interface{}) {
	if err := a.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueryString",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) PutSourceIp(value *AwsAlbListenerRule_SourceIpProperty) {
	if err := a.validatePutSourceIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceIp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) ResetHostHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetHostHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) ResetHttpHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) ResetHttpRequestMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpRequestMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) ResetPathPattern() {
	_jsii_.InvokeVoid(
		a,
		"resetPathPattern",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) ResetSourceIp() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceIp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAlbListenerRule_ConditionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

