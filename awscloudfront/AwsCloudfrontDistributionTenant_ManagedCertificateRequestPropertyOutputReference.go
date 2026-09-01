package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CertificateTransparencyLoggingPreference() *string
	// Experimental.
	SetCertificateTransparencyLoggingPreference(val *string)
	// Experimental.
	CertificateTransparencyLoggingPreferenceInput() *string
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
	PrimaryDomainName() *string
	// Experimental.
	SetPrimaryDomainName(val *string)
	// Experimental.
	PrimaryDomainNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ValidationTokenHost() *string
	// Experimental.
	SetValidationTokenHost(val *string)
	// Experimental.
	ValidationTokenHostInput() *string
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
	ResetCertificateTransparencyLoggingPreference()
	// Experimental.
	ResetPrimaryDomainName()
	// Experimental.
	ResetValidationTokenHost()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference
type jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) CertificateTransparencyLoggingPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateTransparencyLoggingPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) CertificateTransparencyLoggingPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateTransparencyLoggingPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) PrimaryDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) PrimaryDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) ValidationTokenHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationTokenHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) ValidationTokenHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationTokenHostInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsCloudfrontDistributionTenant.ManagedCertificateRequestPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference_Override(a AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsCloudfrontDistributionTenant.ManagedCertificateRequestPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference)SetCertificateTransparencyLoggingPreference(val *string) {
	if err := j.validateSetCertificateTransparencyLoggingPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateTransparencyLoggingPreference",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference)SetPrimaryDomainName(val *string) {
	if err := j.validateSetPrimaryDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryDomainName",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference)SetValidationTokenHost(val *string) {
	if err := j.validateSetValidationTokenHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validationTokenHost",
		val,
	)
}

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) ResetCertificateTransparencyLoggingPreference() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateTransparencyLoggingPreference",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) ResetPrimaryDomainName() {
	_jsii_.InvokeVoid(
		a,
		"resetPrimaryDomainName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) ResetValidationTokenHost() {
	_jsii_.InvokeVoid(
		a,
		"resetValidationTokenHost",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCloudfrontDistributionTenant_ManagedCertificateRequestPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

