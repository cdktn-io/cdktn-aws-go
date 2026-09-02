package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDistribution_ViewerCertificatePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AcmCertificateArn() *string
	// Experimental.
	SetAcmCertificateArn(val *string)
	// Experimental.
	AcmCertificateArnInput() *string
	// Experimental.
	CloudfrontDefaultCertificate() interface{}
	// Experimental.
	SetCloudfrontDefaultCertificate(val interface{})
	// Experimental.
	CloudfrontDefaultCertificateInput() interface{}
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
	IamCertificateId() *string
	// Experimental.
	SetIamCertificateId(val *string)
	// Experimental.
	IamCertificateIdInput() *string
	// Experimental.
	InternalValue() *TfDistribution_ViewerCertificateProperty
	// Experimental.
	SetInternalValue(val *TfDistribution_ViewerCertificateProperty)
	// Experimental.
	MinimumProtocolVersion() *string
	// Experimental.
	SetMinimumProtocolVersion(val *string)
	// Experimental.
	MinimumProtocolVersionInput() *string
	// Experimental.
	SslSupportMethod() *string
	// Experimental.
	SetSslSupportMethod(val *string)
	// Experimental.
	SslSupportMethodInput() *string
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
	ResetAcmCertificateArn()
	// Experimental.
	ResetCloudfrontDefaultCertificate()
	// Experimental.
	ResetIamCertificateId()
	// Experimental.
	ResetMinimumProtocolVersion()
	// Experimental.
	ResetSslSupportMethod()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDistribution_ViewerCertificatePropertyOutputReference
type jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) AcmCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acmCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) AcmCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acmCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) CloudfrontDefaultCertificate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudfrontDefaultCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) CloudfrontDefaultCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudfrontDefaultCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) IamCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) IamCertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamCertificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) InternalValue() *TfDistribution_ViewerCertificateProperty {
	var returns *TfDistribution_ViewerCertificateProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) MinimumProtocolVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumProtocolVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) MinimumProtocolVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumProtocolVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) SslSupportMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslSupportMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) SslSupportMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslSupportMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDistribution_ViewerCertificatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDistribution_ViewerCertificatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDistribution_ViewerCertificatePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfDistribution.ViewerCertificatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDistribution_ViewerCertificatePropertyOutputReference_Override(t TfDistribution_ViewerCertificatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfDistribution.ViewerCertificatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference)SetAcmCertificateArn(val *string) {
	if err := j.validateSetAcmCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acmCertificateArn",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference)SetCloudfrontDefaultCertificate(val interface{}) {
	if err := j.validateSetCloudfrontDefaultCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudfrontDefaultCertificate",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference)SetIamCertificateId(val *string) {
	if err := j.validateSetIamCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamCertificateId",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference)SetInternalValue(val *TfDistribution_ViewerCertificateProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference)SetMinimumProtocolVersion(val *string) {
	if err := j.validateSetMinimumProtocolVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumProtocolVersion",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference)SetSslSupportMethod(val *string) {
	if err := j.validateSetSslSupportMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslSupportMethod",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) ResetAcmCertificateArn() {
	_jsii_.InvokeVoid(
		t,
		"resetAcmCertificateArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) ResetCloudfrontDefaultCertificate() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudfrontDefaultCertificate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) ResetIamCertificateId() {
	_jsii_.InvokeVoid(
		t,
		"resetIamCertificateId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) ResetMinimumProtocolVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetMinimumProtocolVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) ResetSslSupportMethod() {
	_jsii_.InvokeVoid(
		t,
		"resetSslSupportMethod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDistribution_ViewerCertificatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

