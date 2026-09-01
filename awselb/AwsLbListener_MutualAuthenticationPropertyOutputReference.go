package awselb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLbListener_MutualAuthenticationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdvertiseTrustStoreCaNames() *string
	// Experimental.
	SetAdvertiseTrustStoreCaNames(val *string)
	// Experimental.
	AdvertiseTrustStoreCaNamesInput() *string
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
	IgnoreClientCertificateExpiry() interface{}
	// Experimental.
	SetIgnoreClientCertificateExpiry(val interface{})
	// Experimental.
	IgnoreClientCertificateExpiryInput() interface{}
	// Experimental.
	InternalValue() *AwsLbListener_MutualAuthenticationProperty
	// Experimental.
	SetInternalValue(val *AwsLbListener_MutualAuthenticationProperty)
	// Experimental.
	Mode() *string
	// Experimental.
	SetMode(val *string)
	// Experimental.
	ModeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TrustStoreArn() *string
	// Experimental.
	SetTrustStoreArn(val *string)
	// Experimental.
	TrustStoreArnInput() *string
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
	ResetAdvertiseTrustStoreCaNames()
	// Experimental.
	ResetIgnoreClientCertificateExpiry()
	// Experimental.
	ResetTrustStoreArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLbListener_MutualAuthenticationPropertyOutputReference
type jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) AdvertiseTrustStoreCaNames() *string {
	var returns *string
	_jsii_.Get(
		j,
		"advertiseTrustStoreCaNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) AdvertiseTrustStoreCaNamesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"advertiseTrustStoreCaNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) IgnoreClientCertificateExpiry() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreClientCertificateExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) IgnoreClientCertificateExpiryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreClientCertificateExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) InternalValue() *AwsLbListener_MutualAuthenticationProperty {
	var returns *AwsLbListener_MutualAuthenticationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) TrustStoreArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) TrustStoreArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreArnInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLbListener_MutualAuthenticationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsLbListener_MutualAuthenticationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLbListener_MutualAuthenticationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elb.AwsLbListener.MutualAuthenticationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLbListener_MutualAuthenticationPropertyOutputReference_Override(a AwsLbListener_MutualAuthenticationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.AwsLbListener.MutualAuthenticationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference)SetAdvertiseTrustStoreCaNames(val *string) {
	if err := j.validateSetAdvertiseTrustStoreCaNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advertiseTrustStoreCaNames",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference)SetIgnoreClientCertificateExpiry(val interface{}) {
	if err := j.validateSetIgnoreClientCertificateExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreClientCertificateExpiry",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference)SetInternalValue(val *AwsLbListener_MutualAuthenticationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference)SetTrustStoreArn(val *string) {
	if err := j.validateSetTrustStoreArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustStoreArn",
		val,
	)
}

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) ResetAdvertiseTrustStoreCaNames() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvertiseTrustStoreCaNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) ResetIgnoreClientCertificateExpiry() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreClientCertificateExpiry",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) ResetTrustStoreArn() {
	_jsii_.InvokeVoid(
		a,
		"resetTrustStoreArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLbListener_MutualAuthenticationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

