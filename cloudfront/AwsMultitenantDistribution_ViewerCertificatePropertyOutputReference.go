package cloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference interface {
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

// The jsii proxy struct for AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference
type jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) AcmCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acmCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) AcmCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acmCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) CloudfrontDefaultCertificate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudfrontDefaultCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) CloudfrontDefaultCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudfrontDefaultCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) MinimumProtocolVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumProtocolVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) MinimumProtocolVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumProtocolVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) SslSupportMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslSupportMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) SslSupportMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslSupportMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMultitenantDistribution_ViewerCertificatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMultitenantDistribution_ViewerCertificatePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsMultitenantDistribution.ViewerCertificatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMultitenantDistribution_ViewerCertificatePropertyOutputReference_Override(a AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsMultitenantDistribution.ViewerCertificatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetAcmCertificateArn(val *string) {
	if err := j.validateSetAcmCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acmCertificateArn",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetCloudfrontDefaultCertificate(val interface{}) {
	if err := j.validateSetCloudfrontDefaultCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudfrontDefaultCertificate",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetMinimumProtocolVersion(val *string) {
	if err := j.validateSetMinimumProtocolVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumProtocolVersion",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetSslSupportMethod(val *string) {
	if err := j.validateSetSslSupportMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslSupportMethod",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) ResetAcmCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetAcmCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) ResetCloudfrontDefaultCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudfrontDefaultCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) ResetMinimumProtocolVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetMinimumProtocolVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) ResetSslSupportMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetSslSupportMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMultitenantDistribution_ViewerCertificatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

