package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMultitenantDistribution_ViewerCertificatePropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
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

// The jsii proxy struct for TfMultitenantDistribution_ViewerCertificatePropertyOutputReference
type jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) AcmCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acmCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) AcmCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acmCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) CloudfrontDefaultCertificate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudfrontDefaultCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) CloudfrontDefaultCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudfrontDefaultCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) MinimumProtocolVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumProtocolVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) MinimumProtocolVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumProtocolVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) SslSupportMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslSupportMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) SslSupportMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslSupportMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMultitenantDistribution_ViewerCertificatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfMultitenantDistribution_ViewerCertificatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfMultitenantDistribution_ViewerCertificatePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfMultitenantDistribution.ViewerCertificatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMultitenantDistribution_ViewerCertificatePropertyOutputReference_Override(t TfMultitenantDistribution_ViewerCertificatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfMultitenantDistribution.ViewerCertificatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetAcmCertificateArn(val *string) {
	if err := j.validateSetAcmCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acmCertificateArn",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetCloudfrontDefaultCertificate(val interface{}) {
	if err := j.validateSetCloudfrontDefaultCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudfrontDefaultCertificate",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetMinimumProtocolVersion(val *string) {
	if err := j.validateSetMinimumProtocolVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumProtocolVersion",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetSslSupportMethod(val *string) {
	if err := j.validateSetSslSupportMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslSupportMethod",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) ResetAcmCertificateArn() {
	_jsii_.InvokeVoid(
		t,
		"resetAcmCertificateArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) ResetCloudfrontDefaultCertificate() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudfrontDefaultCertificate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) ResetMinimumProtocolVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetMinimumProtocolVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) ResetSslSupportMethod() {
	_jsii_.InvokeVoid(
		t,
		"resetSslSupportMethod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfMultitenantDistribution_ViewerCertificatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

