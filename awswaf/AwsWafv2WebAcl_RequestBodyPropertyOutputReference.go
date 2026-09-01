package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWafv2WebAcl_RequestBodyPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApiGateway() AwsWafv2WebAcl_ApiGatewayPropertyOutputReference
	// Experimental.
	ApiGatewayInput() *AwsWafv2WebAcl_ApiGatewayProperty
	// Experimental.
	AppRunnerService() AwsWafv2WebAcl_AppRunnerServicePropertyOutputReference
	// Experimental.
	AppRunnerServiceInput() *AwsWafv2WebAcl_AppRunnerServiceProperty
	// Experimental.
	Cloudfront() AwsWafv2WebAcl_CloudfrontPropertyOutputReference
	// Experimental.
	CloudfrontInput() *AwsWafv2WebAcl_CloudfrontProperty
	// Experimental.
	CognitoUserPool() AwsWafv2WebAcl_CognitoUserPoolPropertyOutputReference
	// Experimental.
	CognitoUserPoolInput() *AwsWafv2WebAcl_CognitoUserPoolProperty
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
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VerifiedAccessInstance() AwsWafv2WebAcl_VerifiedAccessInstancePropertyOutputReference
	// Experimental.
	VerifiedAccessInstanceInput() *AwsWafv2WebAcl_VerifiedAccessInstanceProperty
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
	PutApiGateway(value *AwsWafv2WebAcl_ApiGatewayProperty)
	// Experimental.
	PutAppRunnerService(value *AwsWafv2WebAcl_AppRunnerServiceProperty)
	// Experimental.
	PutCloudfront(value *AwsWafv2WebAcl_CloudfrontProperty)
	// Experimental.
	PutCognitoUserPool(value *AwsWafv2WebAcl_CognitoUserPoolProperty)
	// Experimental.
	PutVerifiedAccessInstance(value *AwsWafv2WebAcl_VerifiedAccessInstanceProperty)
	// Experimental.
	ResetApiGateway()
	// Experimental.
	ResetAppRunnerService()
	// Experimental.
	ResetCloudfront()
	// Experimental.
	ResetCognitoUserPool()
	// Experimental.
	ResetVerifiedAccessInstance()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsWafv2WebAcl_RequestBodyPropertyOutputReference
type jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) ApiGateway() AwsWafv2WebAcl_ApiGatewayPropertyOutputReference {
	var returns AwsWafv2WebAcl_ApiGatewayPropertyOutputReference
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) ApiGatewayInput() *AwsWafv2WebAcl_ApiGatewayProperty {
	var returns *AwsWafv2WebAcl_ApiGatewayProperty
	_jsii_.Get(
		j,
		"apiGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) AppRunnerService() AwsWafv2WebAcl_AppRunnerServicePropertyOutputReference {
	var returns AwsWafv2WebAcl_AppRunnerServicePropertyOutputReference
	_jsii_.Get(
		j,
		"appRunnerService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) AppRunnerServiceInput() *AwsWafv2WebAcl_AppRunnerServiceProperty {
	var returns *AwsWafv2WebAcl_AppRunnerServiceProperty
	_jsii_.Get(
		j,
		"appRunnerServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) Cloudfront() AwsWafv2WebAcl_CloudfrontPropertyOutputReference {
	var returns AwsWafv2WebAcl_CloudfrontPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudfront",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) CloudfrontInput() *AwsWafv2WebAcl_CloudfrontProperty {
	var returns *AwsWafv2WebAcl_CloudfrontProperty
	_jsii_.Get(
		j,
		"cloudfrontInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) CognitoUserPool() AwsWafv2WebAcl_CognitoUserPoolPropertyOutputReference {
	var returns AwsWafv2WebAcl_CognitoUserPoolPropertyOutputReference
	_jsii_.Get(
		j,
		"cognitoUserPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) CognitoUserPoolInput() *AwsWafv2WebAcl_CognitoUserPoolProperty {
	var returns *AwsWafv2WebAcl_CognitoUserPoolProperty
	_jsii_.Get(
		j,
		"cognitoUserPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) VerifiedAccessInstance() AwsWafv2WebAcl_VerifiedAccessInstancePropertyOutputReference {
	var returns AwsWafv2WebAcl_VerifiedAccessInstancePropertyOutputReference
	_jsii_.Get(
		j,
		"verifiedAccessInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) VerifiedAccessInstanceInput() *AwsWafv2WebAcl_VerifiedAccessInstanceProperty {
	var returns *AwsWafv2WebAcl_VerifiedAccessInstanceProperty
	_jsii_.Get(
		j,
		"verifiedAccessInstanceInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWafv2WebAcl_RequestBodyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsWafv2WebAcl_RequestBodyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWafv2WebAcl_RequestBodyPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWafv2WebAcl.RequestBodyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWafv2WebAcl_RequestBodyPropertyOutputReference_Override(a AwsWafv2WebAcl_RequestBodyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWafv2WebAcl.RequestBodyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) PutApiGateway(value *AwsWafv2WebAcl_ApiGatewayProperty) {
	if err := a.validatePutApiGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApiGateway",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) PutAppRunnerService(value *AwsWafv2WebAcl_AppRunnerServiceProperty) {
	if err := a.validatePutAppRunnerServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAppRunnerService",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) PutCloudfront(value *AwsWafv2WebAcl_CloudfrontProperty) {
	if err := a.validatePutCloudfrontParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudfront",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) PutCognitoUserPool(value *AwsWafv2WebAcl_CognitoUserPoolProperty) {
	if err := a.validatePutCognitoUserPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCognitoUserPool",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) PutVerifiedAccessInstance(value *AwsWafv2WebAcl_VerifiedAccessInstanceProperty) {
	if err := a.validatePutVerifiedAccessInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVerifiedAccessInstance",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) ResetApiGateway() {
	_jsii_.InvokeVoid(
		a,
		"resetApiGateway",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) ResetAppRunnerService() {
	_jsii_.InvokeVoid(
		a,
		"resetAppRunnerService",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) ResetCloudfront() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudfront",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) ResetCognitoUserPool() {
	_jsii_.InvokeVoid(
		a,
		"resetCognitoUserPool",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) ResetVerifiedAccessInstance() {
	_jsii_.InvokeVoid(
		a,
		"resetVerifiedAccessInstance",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWafv2WebAcl_RequestBodyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

