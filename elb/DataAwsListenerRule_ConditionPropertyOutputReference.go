package elb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsListenerRule_ConditionPropertyOutputReference interface {
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
	HostHeader() DataAwsListenerRule_HostHeaderPropertyList
	// Experimental.
	HostHeaderInput() interface{}
	// Experimental.
	HttpHeader() DataAwsListenerRule_HttpHeaderPropertyList
	// Experimental.
	HttpHeaderInput() interface{}
	// Experimental.
	HttpRequestMethod() DataAwsListenerRule_HttpRequestMethodPropertyList
	// Experimental.
	HttpRequestMethodInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PathPattern() DataAwsListenerRule_PathPatternPropertyList
	// Experimental.
	PathPatternInput() interface{}
	// Experimental.
	QueryString() DataAwsListenerRule_QueryStringPropertyList
	// Experimental.
	QueryStringInput() interface{}
	// Experimental.
	SourceIp() DataAwsListenerRule_SourceIpPropertyList
	// Experimental.
	SourceIpInput() interface{}
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
	PutHostHeader(value interface{})
	// Experimental.
	PutHttpHeader(value interface{})
	// Experimental.
	PutHttpRequestMethod(value interface{})
	// Experimental.
	PutPathPattern(value interface{})
	// Experimental.
	PutQueryString(value interface{})
	// Experimental.
	PutSourceIp(value interface{})
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

// The jsii proxy struct for DataAwsListenerRule_ConditionPropertyOutputReference
type jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) HostHeader() DataAwsListenerRule_HostHeaderPropertyList {
	var returns DataAwsListenerRule_HostHeaderPropertyList
	_jsii_.Get(
		j,
		"hostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) HostHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) HttpHeader() DataAwsListenerRule_HttpHeaderPropertyList {
	var returns DataAwsListenerRule_HttpHeaderPropertyList
	_jsii_.Get(
		j,
		"httpHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) HttpHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) HttpRequestMethod() DataAwsListenerRule_HttpRequestMethodPropertyList {
	var returns DataAwsListenerRule_HttpRequestMethodPropertyList
	_jsii_.Get(
		j,
		"httpRequestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) HttpRequestMethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpRequestMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) PathPattern() DataAwsListenerRule_PathPatternPropertyList {
	var returns DataAwsListenerRule_PathPatternPropertyList
	_jsii_.Get(
		j,
		"pathPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) PathPatternInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pathPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) QueryString() DataAwsListenerRule_QueryStringPropertyList {
	var returns DataAwsListenerRule_QueryStringPropertyList
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) SourceIp() DataAwsListenerRule_SourceIpPropertyList {
	var returns DataAwsListenerRule_SourceIpPropertyList
	_jsii_.Get(
		j,
		"sourceIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) SourceIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsListenerRule_ConditionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsListenerRule_ConditionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsListenerRule_ConditionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elb.DataAwsListenerRule.ConditionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsListenerRule_ConditionPropertyOutputReference_Override(d DataAwsListenerRule_ConditionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.DataAwsListenerRule.ConditionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) PutHostHeader(value interface{}) {
	if err := d.validatePutHostHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHostHeader",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) PutHttpHeader(value interface{}) {
	if err := d.validatePutHttpHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHttpHeader",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) PutHttpRequestMethod(value interface{}) {
	if err := d.validatePutHttpRequestMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHttpRequestMethod",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) PutPathPattern(value interface{}) {
	if err := d.validatePutPathPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPathPattern",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) PutQueryString(value interface{}) {
	if err := d.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueryString",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) PutSourceIp(value interface{}) {
	if err := d.validatePutSourceIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSourceIp",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) ResetHostHeader() {
	_jsii_.InvokeVoid(
		d,
		"resetHostHeader",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) ResetHttpHeader() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpHeader",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) ResetHttpRequestMethod() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpRequestMethod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) ResetPathPattern() {
	_jsii_.InvokeVoid(
		d,
		"resetPathPattern",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryString",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) ResetSourceIp() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceIp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsListenerRule_ConditionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

