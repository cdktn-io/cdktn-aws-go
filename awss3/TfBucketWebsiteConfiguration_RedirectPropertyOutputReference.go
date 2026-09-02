package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfBucketWebsiteConfiguration_RedirectPropertyOutputReference interface {
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
	HostName() *string
	// Experimental.
	SetHostName(val *string)
	// Experimental.
	HostNameInput() *string
	// Experimental.
	HttpRedirectCode() *string
	// Experimental.
	SetHttpRedirectCode(val *string)
	// Experimental.
	HttpRedirectCodeInput() *string
	// Experimental.
	InternalValue() *TfBucketWebsiteConfiguration_RedirectProperty
	// Experimental.
	SetInternalValue(val *TfBucketWebsiteConfiguration_RedirectProperty)
	// Experimental.
	Protocol() *string
	// Experimental.
	SetProtocol(val *string)
	// Experimental.
	ProtocolInput() *string
	// Experimental.
	ReplaceKeyPrefixWith() *string
	// Experimental.
	SetReplaceKeyPrefixWith(val *string)
	// Experimental.
	ReplaceKeyPrefixWithInput() *string
	// Experimental.
	ReplaceKeyWith() *string
	// Experimental.
	SetReplaceKeyWith(val *string)
	// Experimental.
	ReplaceKeyWithInput() *string
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
	ResetHostName()
	// Experimental.
	ResetHttpRedirectCode()
	// Experimental.
	ResetProtocol()
	// Experimental.
	ResetReplaceKeyPrefixWith()
	// Experimental.
	ResetReplaceKeyWith()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfBucketWebsiteConfiguration_RedirectPropertyOutputReference
type jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) HttpRedirectCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRedirectCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) HttpRedirectCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRedirectCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) InternalValue() *TfBucketWebsiteConfiguration_RedirectProperty {
	var returns *TfBucketWebsiteConfiguration_RedirectProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) ReplaceKeyPrefixWith() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceKeyPrefixWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) ReplaceKeyPrefixWithInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceKeyPrefixWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) ReplaceKeyWith() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceKeyWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) ReplaceKeyWithInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceKeyWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfBucketWebsiteConfiguration_RedirectPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfBucketWebsiteConfiguration_RedirectPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfBucketWebsiteConfiguration_RedirectPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.TfBucketWebsiteConfiguration.RedirectPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfBucketWebsiteConfiguration_RedirectPropertyOutputReference_Override(t TfBucketWebsiteConfiguration_RedirectPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.TfBucketWebsiteConfiguration.RedirectPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference)SetHostName(val *string) {
	if err := j.validateSetHostNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference)SetHttpRedirectCode(val *string) {
	if err := j.validateSetHttpRedirectCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpRedirectCode",
		val,
	)
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference)SetInternalValue(val *TfBucketWebsiteConfiguration_RedirectProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference)SetReplaceKeyPrefixWith(val *string) {
	if err := j.validateSetReplaceKeyPrefixWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceKeyPrefixWith",
		val,
	)
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference)SetReplaceKeyWith(val *string) {
	if err := j.validateSetReplaceKeyWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceKeyWith",
		val,
	)
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) ResetHostName() {
	_jsii_.InvokeVoid(
		t,
		"resetHostName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) ResetHttpRedirectCode() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpRedirectCode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		t,
		"resetProtocol",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) ResetReplaceKeyPrefixWith() {
	_jsii_.InvokeVoid(
		t,
		"resetReplaceKeyPrefixWith",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) ResetReplaceKeyWith() {
	_jsii_.InvokeVoid(
		t,
		"resetReplaceKeyWith",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfBucketWebsiteConfiguration_RedirectPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

