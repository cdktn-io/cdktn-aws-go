package cloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMultitenantDistribution_OriginPropertyOutputReference interface {
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
	CustomHeader() AwsMultitenantDistribution_CustomHeaderPropertyList
	// Experimental.
	CustomHeaderInput() interface{}
	// Experimental.
	CustomOriginConfig() AwsMultitenantDistribution_CustomOriginConfigPropertyList
	// Experimental.
	CustomOriginConfigInput() interface{}
	// Experimental.
	DomainName() *string
	// Experimental.
	SetDomainName(val *string)
	// Experimental.
	DomainNameInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
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
	OriginPath() *string
	// Experimental.
	SetOriginPath(val *string)
	// Experimental.
	OriginPathInput() *string
	// Experimental.
	OriginShield() AwsMultitenantDistribution_OriginShieldPropertyList
	// Experimental.
	OriginShieldInput() interface{}
	// Experimental.
	ResponseCompletionTimeout() *float64
	// Experimental.
	SetResponseCompletionTimeout(val *float64)
	// Experimental.
	ResponseCompletionTimeoutInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VpcOriginConfig() AwsMultitenantDistribution_VpcOriginConfigPropertyList
	// Experimental.
	VpcOriginConfigInput() interface{}
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
	PutCustomOriginConfig(value interface{})
	// Experimental.
	PutOriginShield(value interface{})
	// Experimental.
	PutVpcOriginConfig(value interface{})
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

// The jsii proxy struct for AwsMultitenantDistribution_OriginPropertyOutputReference
type jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ConnectionAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ConnectionAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ConnectionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ConnectionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) CustomHeader() AwsMultitenantDistribution_CustomHeaderPropertyList {
	var returns AwsMultitenantDistribution_CustomHeaderPropertyList
	_jsii_.Get(
		j,
		"customHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) CustomHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) CustomOriginConfig() AwsMultitenantDistribution_CustomOriginConfigPropertyList {
	var returns AwsMultitenantDistribution_CustomOriginConfigPropertyList
	_jsii_.Get(
		j,
		"customOriginConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) CustomOriginConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customOriginConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) OriginAccessControlId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessControlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) OriginAccessControlIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessControlIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) OriginPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) OriginPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) OriginShield() AwsMultitenantDistribution_OriginShieldPropertyList {
	var returns AwsMultitenantDistribution_OriginShieldPropertyList
	_jsii_.Get(
		j,
		"originShield",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) OriginShieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originShieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ResponseCompletionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"responseCompletionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ResponseCompletionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"responseCompletionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) VpcOriginConfig() AwsMultitenantDistribution_VpcOriginConfigPropertyList {
	var returns AwsMultitenantDistribution_VpcOriginConfigPropertyList
	_jsii_.Get(
		j,
		"vpcOriginConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) VpcOriginConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcOriginConfigInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMultitenantDistribution_OriginPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsMultitenantDistribution_OriginPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMultitenantDistribution_OriginPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsMultitenantDistribution.OriginPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMultitenantDistribution_OriginPropertyOutputReference_Override(a AwsMultitenantDistribution_OriginPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsMultitenantDistribution.OriginPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference)SetConnectionAttempts(val *float64) {
	if err := j.validateSetConnectionAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionAttempts",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference)SetConnectionTimeout(val *float64) {
	if err := j.validateSetConnectionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference)SetOriginAccessControlId(val *string) {
	if err := j.validateSetOriginAccessControlIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originAccessControlId",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference)SetOriginPath(val *string) {
	if err := j.validateSetOriginPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originPath",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference)SetResponseCompletionTimeout(val *float64) {
	if err := j.validateSetResponseCompletionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseCompletionTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) PutCustomHeader(value interface{}) {
	if err := a.validatePutCustomHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomHeader",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) PutCustomOriginConfig(value interface{}) {
	if err := a.validatePutCustomOriginConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomOriginConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) PutOriginShield(value interface{}) {
	if err := a.validatePutOriginShieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOriginShield",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) PutVpcOriginConfig(value interface{}) {
	if err := a.validatePutVpcOriginConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcOriginConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ResetConnectionAttempts() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionAttempts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ResetConnectionTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ResetCustomHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ResetCustomOriginConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomOriginConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ResetOriginAccessControlId() {
	_jsii_.InvokeVoid(
		a,
		"resetOriginAccessControlId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ResetOriginPath() {
	_jsii_.InvokeVoid(
		a,
		"resetOriginPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ResetOriginShield() {
	_jsii_.InvokeVoid(
		a,
		"resetOriginShield",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ResetResponseCompletionTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseCompletionTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ResetVpcOriginConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcOriginConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMultitenantDistribution_OriginPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

