package awsacmpca

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsacmpca/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsacmpca/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCertificateAuthority_CrlConfigurationPropertyOutputReference interface {
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
	InternalValue() *TfCertificateAuthority_CrlConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfCertificateAuthority_CrlConfigurationProperty)
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

// The jsii proxy struct for TfCertificateAuthority_CrlConfigurationPropertyOutputReference
type jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) CustomCname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) CustomCnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) CustomPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) CustomPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) ExpirationInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) ExpirationInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) InternalValue() *TfCertificateAuthority_CrlConfigurationProperty {
	var returns *TfCertificateAuthority_CrlConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) S3BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) S3ObjectAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) S3ObjectAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCertificateAuthority_CrlConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCertificateAuthority_CrlConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCertificateAuthority_CrlConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-acm-pca.TfCertificateAuthority.CrlConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCertificateAuthority_CrlConfigurationPropertyOutputReference_Override(t TfCertificateAuthority_CrlConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-acm-pca.TfCertificateAuthority.CrlConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference)SetCustomCname(val *string) {
	if err := j.validateSetCustomCnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customCname",
		val,
	)
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference)SetCustomPath(val *string) {
	if err := j.validateSetCustomPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customPath",
		val,
	)
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference)SetExpirationInDays(val *float64) {
	if err := j.validateSetExpirationInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationInDays",
		val,
	)
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference)SetInternalValue(val *TfCertificateAuthority_CrlConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference)SetS3BucketName(val *string) {
	if err := j.validateSetS3BucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference)SetS3ObjectAcl(val *string) {
	if err := j.validateSetS3ObjectAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3ObjectAcl",
		val,
	)
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) ResetCustomCname() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomCname",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) ResetCustomPath() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) ResetExpirationInDays() {
	_jsii_.InvokeVoid(
		t,
		"resetExpirationInDays",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) ResetS3BucketName() {
	_jsii_.InvokeVoid(
		t,
		"resetS3BucketName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) ResetS3ObjectAcl() {
	_jsii_.InvokeVoid(
		t,
		"resetS3ObjectAcl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCertificateAuthority_CrlConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

