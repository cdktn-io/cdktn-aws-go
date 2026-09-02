package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDistribution_OriginPropertyOutputReference interface {
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
	// Experimental.
	ConnectionAttempts() *float64
	// Experimental.
	SetConnectionAttempts(val *float64)
	// Experimental.
	ConnectionAttemptsInput() *float64
	// Experimental.
	ConnectionTimeout() *float64
	// Experimental.
	SetConnectionTimeout(val *float64)
	// Experimental.
	ConnectionTimeoutInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CustomHeader() TfDistribution_CustomHeaderPropertyList
	// Experimental.
	CustomHeaderInput() interface{}
	// Experimental.
	CustomOriginConfig() TfDistribution_CustomOriginConfigPropertyOutputReference
	// Experimental.
	CustomOriginConfigInput() *TfDistribution_CustomOriginConfigProperty
	// Experimental.
	DomainName() *string
	// Experimental.
	SetDomainName(val *string)
	// Experimental.
	DomainNameInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	OriginAccessControlId() *string
	// Experimental.
	SetOriginAccessControlId(val *string)
	// Experimental.
	OriginAccessControlIdInput() *string
	// Experimental.
	OriginId() *string
	// Experimental.
	SetOriginId(val *string)
	// Experimental.
	OriginIdInput() *string
	// Experimental.
	OriginPath() *string
	// Experimental.
	SetOriginPath(val *string)
	// Experimental.
	OriginPathInput() *string
	// Experimental.
	OriginShield() TfDistribution_OriginShieldPropertyOutputReference
	// Experimental.
	OriginShieldInput() *TfDistribution_OriginShieldProperty
	// Experimental.
	ResponseCompletionTimeout() *float64
	// Experimental.
	SetResponseCompletionTimeout(val *float64)
	// Experimental.
	ResponseCompletionTimeoutInput() *float64
	// Experimental.
	S3OriginConfig() TfDistribution_S3OriginConfigPropertyOutputReference
	// Experimental.
	S3OriginConfigInput() *TfDistribution_S3OriginConfigProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VpcOriginConfig() TfDistribution_VpcOriginConfigPropertyOutputReference
	// Experimental.
	VpcOriginConfigInput() *TfDistribution_VpcOriginConfigProperty
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
	PutCustomHeader(value interface{})
	// Experimental.
	PutCustomOriginConfig(value *TfDistribution_CustomOriginConfigProperty)
	// Experimental.
	PutOriginShield(value *TfDistribution_OriginShieldProperty)
	// Experimental.
	PutS3OriginConfig(value *TfDistribution_S3OriginConfigProperty)
	// Experimental.
	PutVpcOriginConfig(value *TfDistribution_VpcOriginConfigProperty)
	// Experimental.
	ResetConnectionAttempts()
	// Experimental.
	ResetConnectionTimeout()
	// Experimental.
	ResetCustomHeader()
	// Experimental.
	ResetCustomOriginConfig()
	// Experimental.
	ResetOriginAccessControlId()
	// Experimental.
	ResetOriginPath()
	// Experimental.
	ResetOriginShield()
	// Experimental.
	ResetResponseCompletionTimeout()
	// Experimental.
	ResetS3OriginConfig()
	// Experimental.
	ResetVpcOriginConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDistribution_OriginPropertyOutputReference
type jsiiProxy_TfDistribution_OriginPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ConnectionAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ConnectionAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ConnectionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ConnectionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) CustomHeader() TfDistribution_CustomHeaderPropertyList {
	var returns TfDistribution_CustomHeaderPropertyList
	_jsii_.Get(
		j,
		"customHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) CustomHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) CustomOriginConfig() TfDistribution_CustomOriginConfigPropertyOutputReference {
	var returns TfDistribution_CustomOriginConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"customOriginConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) CustomOriginConfigInput() *TfDistribution_CustomOriginConfigProperty {
	var returns *TfDistribution_CustomOriginConfigProperty
	_jsii_.Get(
		j,
		"customOriginConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) OriginAccessControlId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessControlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) OriginAccessControlIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessControlIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) OriginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) OriginIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) OriginPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) OriginPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) OriginShield() TfDistribution_OriginShieldPropertyOutputReference {
	var returns TfDistribution_OriginShieldPropertyOutputReference
	_jsii_.Get(
		j,
		"originShield",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) OriginShieldInput() *TfDistribution_OriginShieldProperty {
	var returns *TfDistribution_OriginShieldProperty
	_jsii_.Get(
		j,
		"originShieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ResponseCompletionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"responseCompletionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ResponseCompletionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"responseCompletionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) S3OriginConfig() TfDistribution_S3OriginConfigPropertyOutputReference {
	var returns TfDistribution_S3OriginConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"s3OriginConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) S3OriginConfigInput() *TfDistribution_S3OriginConfigProperty {
	var returns *TfDistribution_S3OriginConfigProperty
	_jsii_.Get(
		j,
		"s3OriginConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) VpcOriginConfig() TfDistribution_VpcOriginConfigPropertyOutputReference {
	var returns TfDistribution_VpcOriginConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcOriginConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference) VpcOriginConfigInput() *TfDistribution_VpcOriginConfigProperty {
	var returns *TfDistribution_VpcOriginConfigProperty
	_jsii_.Get(
		j,
		"vpcOriginConfigInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDistribution_OriginPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfDistribution_OriginPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDistribution_OriginPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDistribution_OriginPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfDistribution.OriginPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDistribution_OriginPropertyOutputReference_Override(t TfDistribution_OriginPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfDistribution.OriginPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference)SetConnectionAttempts(val *float64) {
	if err := j.validateSetConnectionAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionAttempts",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference)SetConnectionTimeout(val *float64) {
	if err := j.validateSetConnectionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionTimeout",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference)SetOriginAccessControlId(val *string) {
	if err := j.validateSetOriginAccessControlIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originAccessControlId",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference)SetOriginId(val *string) {
	if err := j.validateSetOriginIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originId",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference)SetOriginPath(val *string) {
	if err := j.validateSetOriginPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originPath",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference)SetResponseCompletionTimeout(val *float64) {
	if err := j.validateSetResponseCompletionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseCompletionTimeout",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OriginPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) PutCustomHeader(value interface{}) {
	if err := t.validatePutCustomHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) PutCustomOriginConfig(value *TfDistribution_CustomOriginConfigProperty) {
	if err := t.validatePutCustomOriginConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomOriginConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) PutOriginShield(value *TfDistribution_OriginShieldProperty) {
	if err := t.validatePutOriginShieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOriginShield",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) PutS3OriginConfig(value *TfDistribution_S3OriginConfigProperty) {
	if err := t.validatePutS3OriginConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3OriginConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) PutVpcOriginConfig(value *TfDistribution_VpcOriginConfigProperty) {
	if err := t.validatePutVpcOriginConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVpcOriginConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ResetConnectionAttempts() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectionAttempts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ResetConnectionTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectionTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ResetCustomHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ResetCustomOriginConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomOriginConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ResetOriginAccessControlId() {
	_jsii_.InvokeVoid(
		t,
		"resetOriginAccessControlId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ResetOriginPath() {
	_jsii_.InvokeVoid(
		t,
		"resetOriginPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ResetOriginShield() {
	_jsii_.InvokeVoid(
		t,
		"resetOriginShield",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ResetResponseCompletionTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetResponseCompletionTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ResetS3OriginConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetS3OriginConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ResetVpcOriginConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcOriginConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDistribution_OriginPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

