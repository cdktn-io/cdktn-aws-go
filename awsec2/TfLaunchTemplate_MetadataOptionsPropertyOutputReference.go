package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfLaunchTemplate_MetadataOptionsPropertyOutputReference interface {
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
	HttpEndpoint() *string
	// Experimental.
	SetHttpEndpoint(val *string)
	// Experimental.
	HttpEndpointInput() *string
	// Experimental.
	HttpProtocolIpv6() *string
	// Experimental.
	SetHttpProtocolIpv6(val *string)
	// Experimental.
	HttpProtocolIpv6Input() *string
	// Experimental.
	HttpPutResponseHopLimit() *float64
	// Experimental.
	SetHttpPutResponseHopLimit(val *float64)
	// Experimental.
	HttpPutResponseHopLimitInput() *float64
	// Experimental.
	HttpTokens() *string
	// Experimental.
	SetHttpTokens(val *string)
	// Experimental.
	HttpTokensInput() *string
	// Experimental.
	InstanceMetadataTags() *string
	// Experimental.
	SetInstanceMetadataTags(val *string)
	// Experimental.
	InstanceMetadataTagsInput() *string
	// Experimental.
	InternalValue() *TfLaunchTemplate_MetadataOptionsProperty
	// Experimental.
	SetInternalValue(val *TfLaunchTemplate_MetadataOptionsProperty)
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
	ResetHttpEndpoint()
	// Experimental.
	ResetHttpProtocolIpv6()
	// Experimental.
	ResetHttpPutResponseHopLimit()
	// Experimental.
	ResetHttpTokens()
	// Experimental.
	ResetInstanceMetadataTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfLaunchTemplate_MetadataOptionsPropertyOutputReference
type jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) HttpEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) HttpEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) HttpProtocolIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProtocolIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) HttpProtocolIpv6Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProtocolIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) HttpPutResponseHopLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPutResponseHopLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) HttpPutResponseHopLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPutResponseHopLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) HttpTokens() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) HttpTokensInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) InstanceMetadataTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceMetadataTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) InstanceMetadataTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceMetadataTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) InternalValue() *TfLaunchTemplate_MetadataOptionsProperty {
	var returns *TfLaunchTemplate_MetadataOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfLaunchTemplate_MetadataOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfLaunchTemplate_MetadataOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfLaunchTemplate_MetadataOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2.TfLaunchTemplate.MetadataOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfLaunchTemplate_MetadataOptionsPropertyOutputReference_Override(t TfLaunchTemplate_MetadataOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.TfLaunchTemplate.MetadataOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference)SetHttpEndpoint(val *string) {
	if err := j.validateSetHttpEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpEndpoint",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference)SetHttpProtocolIpv6(val *string) {
	if err := j.validateSetHttpProtocolIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpProtocolIpv6",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference)SetHttpPutResponseHopLimit(val *float64) {
	if err := j.validateSetHttpPutResponseHopLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpPutResponseHopLimit",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference)SetHttpTokens(val *string) {
	if err := j.validateSetHttpTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpTokens",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference)SetInstanceMetadataTags(val *string) {
	if err := j.validateSetInstanceMetadataTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceMetadataTags",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference)SetInternalValue(val *TfLaunchTemplate_MetadataOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) ResetHttpEndpoint() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpEndpoint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) ResetHttpProtocolIpv6() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpProtocolIpv6",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) ResetHttpPutResponseHopLimit() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpPutResponseHopLimit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) ResetHttpTokens() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpTokens",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) ResetInstanceMetadataTags() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceMetadataTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfLaunchTemplate_MetadataOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

