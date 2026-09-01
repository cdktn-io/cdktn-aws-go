package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCloudfrontDistribution_OriginPropertyOutputReference interface {
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
	CustomHeader() AwsCloudfrontDistribution_CustomHeaderPropertyList
	// Experimental.
	CustomHeaderInput() interface{}
	// Experimental.
	CustomOriginConfig() AwsCloudfrontDistribution_CustomOriginConfigPropertyOutputReference
	// Experimental.
	CustomOriginConfigInput() *AwsCloudfrontDistribution_CustomOriginConfigProperty
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
	OriginShield() AwsCloudfrontDistribution_OriginShieldPropertyOutputReference
	// Experimental.
	OriginShieldInput() *AwsCloudfrontDistribution_OriginShieldProperty
	// Experimental.
	ResponseCompletionTimeout() *float64
	// Experimental.
	SetResponseCompletionTimeout(val *float64)
	// Experimental.
	ResponseCompletionTimeoutInput() *float64
	// Experimental.
	S3OriginConfig() AwsCloudfrontDistribution_S3OriginConfigPropertyOutputReference
	// Experimental.
	S3OriginConfigInput() *AwsCloudfrontDistribution_S3OriginConfigProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VpcOriginConfig() AwsCloudfrontDistribution_VpcOriginConfigPropertyOutputReference
	// Experimental.
	VpcOriginConfigInput() *AwsCloudfrontDistribution_VpcOriginConfigProperty
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
	PutCustomOriginConfig(value *AwsCloudfrontDistribution_CustomOriginConfigProperty)
	// Experimental.
	PutOriginShield(value *AwsCloudfrontDistribution_OriginShieldProperty)
	// Experimental.
	PutS3OriginConfig(value *AwsCloudfrontDistribution_S3OriginConfigProperty)
	// Experimental.
	PutVpcOriginConfig(value *AwsCloudfrontDistribution_VpcOriginConfigProperty)
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

// The jsii proxy struct for AwsCloudfrontDistribution_OriginPropertyOutputReference
type jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ConnectionAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ConnectionAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ConnectionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ConnectionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) CustomHeader() AwsCloudfrontDistribution_CustomHeaderPropertyList {
	var returns AwsCloudfrontDistribution_CustomHeaderPropertyList
	_jsii_.Get(
		j,
		"customHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) CustomHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) CustomOriginConfig() AwsCloudfrontDistribution_CustomOriginConfigPropertyOutputReference {
	var returns AwsCloudfrontDistribution_CustomOriginConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"customOriginConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) CustomOriginConfigInput() *AwsCloudfrontDistribution_CustomOriginConfigProperty {
	var returns *AwsCloudfrontDistribution_CustomOriginConfigProperty
	_jsii_.Get(
		j,
		"customOriginConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) OriginAccessControlId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessControlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) OriginAccessControlIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessControlIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) OriginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) OriginIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) OriginPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) OriginPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) OriginShield() AwsCloudfrontDistribution_OriginShieldPropertyOutputReference {
	var returns AwsCloudfrontDistribution_OriginShieldPropertyOutputReference
	_jsii_.Get(
		j,
		"originShield",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) OriginShieldInput() *AwsCloudfrontDistribution_OriginShieldProperty {
	var returns *AwsCloudfrontDistribution_OriginShieldProperty
	_jsii_.Get(
		j,
		"originShieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ResponseCompletionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"responseCompletionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ResponseCompletionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"responseCompletionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) S3OriginConfig() AwsCloudfrontDistribution_S3OriginConfigPropertyOutputReference {
	var returns AwsCloudfrontDistribution_S3OriginConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"s3OriginConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) S3OriginConfigInput() *AwsCloudfrontDistribution_S3OriginConfigProperty {
	var returns *AwsCloudfrontDistribution_S3OriginConfigProperty
	_jsii_.Get(
		j,
		"s3OriginConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) VpcOriginConfig() AwsCloudfrontDistribution_VpcOriginConfigPropertyOutputReference {
	var returns AwsCloudfrontDistribution_VpcOriginConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcOriginConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) VpcOriginConfigInput() *AwsCloudfrontDistribution_VpcOriginConfigProperty {
	var returns *AwsCloudfrontDistribution_VpcOriginConfigProperty
	_jsii_.Get(
		j,
		"vpcOriginConfigInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCloudfrontDistribution_OriginPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsCloudfrontDistribution_OriginPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCloudfrontDistribution_OriginPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsCloudfrontDistribution.OriginPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCloudfrontDistribution_OriginPropertyOutputReference_Override(a AwsCloudfrontDistribution_OriginPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsCloudfrontDistribution.OriginPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference)SetConnectionAttempts(val *float64) {
	if err := j.validateSetConnectionAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionAttempts",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference)SetConnectionTimeout(val *float64) {
	if err := j.validateSetConnectionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference)SetOriginAccessControlId(val *string) {
	if err := j.validateSetOriginAccessControlIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originAccessControlId",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference)SetOriginId(val *string) {
	if err := j.validateSetOriginIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originId",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference)SetOriginPath(val *string) {
	if err := j.validateSetOriginPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originPath",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference)SetResponseCompletionTimeout(val *float64) {
	if err := j.validateSetResponseCompletionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseCompletionTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) PutCustomHeader(value interface{}) {
	if err := a.validatePutCustomHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomHeader",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) PutCustomOriginConfig(value *AwsCloudfrontDistribution_CustomOriginConfigProperty) {
	if err := a.validatePutCustomOriginConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomOriginConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) PutOriginShield(value *AwsCloudfrontDistribution_OriginShieldProperty) {
	if err := a.validatePutOriginShieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOriginShield",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) PutS3OriginConfig(value *AwsCloudfrontDistribution_S3OriginConfigProperty) {
	if err := a.validatePutS3OriginConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3OriginConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) PutVpcOriginConfig(value *AwsCloudfrontDistribution_VpcOriginConfigProperty) {
	if err := a.validatePutVpcOriginConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcOriginConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ResetConnectionAttempts() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionAttempts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ResetConnectionTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ResetCustomHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ResetCustomOriginConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomOriginConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ResetOriginAccessControlId() {
	_jsii_.InvokeVoid(
		a,
		"resetOriginAccessControlId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ResetOriginPath() {
	_jsii_.InvokeVoid(
		a,
		"resetOriginPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ResetOriginShield() {
	_jsii_.InvokeVoid(
		a,
		"resetOriginShield",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ResetResponseCompletionTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseCompletionTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ResetS3OriginConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetS3OriginConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ResetVpcOriginConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcOriginConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCloudfrontDistribution_OriginPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

