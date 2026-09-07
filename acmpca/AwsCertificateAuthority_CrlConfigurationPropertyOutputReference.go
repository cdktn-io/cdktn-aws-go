package acmpca

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/acmpca/jsii"

	"github.com/cdktn-io/cdktn-aws-go/acmpca/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCertificateAuthority_CrlConfigurationPropertyOutputReference interface {
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
	CustomCname() *string
	// Experimental.
	SetCustomCname(val *string)
	// Experimental.
	CustomCnameInput() *string
	// Experimental.
	CustomPath() *string
	// Experimental.
	SetCustomPath(val *string)
	// Experimental.
	CustomPathInput() *string
	// Experimental.
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	ExpirationInDays() *float64
	// Experimental.
	SetExpirationInDays(val *float64)
	// Experimental.
	ExpirationInDaysInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsCertificateAuthority_CrlConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsCertificateAuthority_CrlConfigurationProperty)
	// Experimental.
	S3BucketName() *string
	// Experimental.
	SetS3BucketName(val *string)
	// Experimental.
	S3BucketNameInput() *string
	// Experimental.
	S3ObjectAcl() *string
	// Experimental.
	SetS3ObjectAcl(val *string)
	// Experimental.
	S3ObjectAclInput() *string
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
	ResetCustomCname()
	// Experimental.
	ResetCustomPath()
	// Experimental.
	ResetEnabled()
	// Experimental.
	ResetExpirationInDays()
	// Experimental.
	ResetS3BucketName()
	// Experimental.
	ResetS3ObjectAcl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCertificateAuthority_CrlConfigurationPropertyOutputReference
type jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) CustomCname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) CustomCnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) CustomPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) CustomPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) ExpirationInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) ExpirationInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) InternalValue() *AwsCertificateAuthority_CrlConfigurationProperty {
	var returns *AwsCertificateAuthority_CrlConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) S3BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) S3ObjectAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) S3ObjectAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCertificateAuthority_CrlConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCertificateAuthority_CrlConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCertificateAuthority_CrlConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-acm-pca.AwsCertificateAuthority.CrlConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCertificateAuthority_CrlConfigurationPropertyOutputReference_Override(a AwsCertificateAuthority_CrlConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-acm-pca.AwsCertificateAuthority.CrlConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference)SetCustomCname(val *string) {
	if err := j.validateSetCustomCnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customCname",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference)SetCustomPath(val *string) {
	if err := j.validateSetCustomPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customPath",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference)SetExpirationInDays(val *float64) {
	if err := j.validateSetExpirationInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationInDays",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference)SetInternalValue(val *AwsCertificateAuthority_CrlConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference)SetS3BucketName(val *string) {
	if err := j.validateSetS3BucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference)SetS3ObjectAcl(val *string) {
	if err := j.validateSetS3ObjectAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3ObjectAcl",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) ResetCustomCname() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomCname",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) ResetCustomPath() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) ResetExpirationInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetExpirationInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) ResetS3BucketName() {
	_jsii_.InvokeVoid(
		a,
		"resetS3BucketName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) ResetS3ObjectAcl() {
	_jsii_.InvokeVoid(
		a,
		"resetS3ObjectAcl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCertificateAuthority_CrlConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

