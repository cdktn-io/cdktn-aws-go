package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAcl_RequestBodyPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApiGateway() TfWebAcl_ApiGatewayPropertyOutputReference
	// Experimental.
	ApiGatewayInput() *TfWebAcl_ApiGatewayProperty
	// Experimental.
	AppRunnerService() TfWebAcl_AppRunnerServicePropertyOutputReference
	// Experimental.
	AppRunnerServiceInput() *TfWebAcl_AppRunnerServiceProperty
	// Experimental.
	Cloudfront() TfWebAcl_CloudfrontPropertyOutputReference
	// Experimental.
	CloudfrontInput() *TfWebAcl_CloudfrontProperty
	// Experimental.
	CognitoUserPool() TfWebAcl_CognitoUserPoolPropertyOutputReference
	// Experimental.
	CognitoUserPoolInput() *TfWebAcl_CognitoUserPoolProperty
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
	VerifiedAccessInstance() TfWebAcl_VerifiedAccessInstancePropertyOutputReference
	// Experimental.
	VerifiedAccessInstanceInput() *TfWebAcl_VerifiedAccessInstanceProperty
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
	PutApiGateway(value *TfWebAcl_ApiGatewayProperty)
	// Experimental.
	PutAppRunnerService(value *TfWebAcl_AppRunnerServiceProperty)
	// Experimental.
	PutCloudfront(value *TfWebAcl_CloudfrontProperty)
	// Experimental.
	PutCognitoUserPool(value *TfWebAcl_CognitoUserPoolProperty)
	// Experimental.
	PutVerifiedAccessInstance(value *TfWebAcl_VerifiedAccessInstanceProperty)
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

// The jsii proxy struct for TfWebAcl_RequestBodyPropertyOutputReference
type jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) ApiGateway() TfWebAcl_ApiGatewayPropertyOutputReference {
	var returns TfWebAcl_ApiGatewayPropertyOutputReference
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) ApiGatewayInput() *TfWebAcl_ApiGatewayProperty {
	var returns *TfWebAcl_ApiGatewayProperty
	_jsii_.Get(
		j,
		"apiGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) AppRunnerService() TfWebAcl_AppRunnerServicePropertyOutputReference {
	var returns TfWebAcl_AppRunnerServicePropertyOutputReference
	_jsii_.Get(
		j,
		"appRunnerService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) AppRunnerServiceInput() *TfWebAcl_AppRunnerServiceProperty {
	var returns *TfWebAcl_AppRunnerServiceProperty
	_jsii_.Get(
		j,
		"appRunnerServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) Cloudfront() TfWebAcl_CloudfrontPropertyOutputReference {
	var returns TfWebAcl_CloudfrontPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudfront",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) CloudfrontInput() *TfWebAcl_CloudfrontProperty {
	var returns *TfWebAcl_CloudfrontProperty
	_jsii_.Get(
		j,
		"cloudfrontInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) CognitoUserPool() TfWebAcl_CognitoUserPoolPropertyOutputReference {
	var returns TfWebAcl_CognitoUserPoolPropertyOutputReference
	_jsii_.Get(
		j,
		"cognitoUserPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) CognitoUserPoolInput() *TfWebAcl_CognitoUserPoolProperty {
	var returns *TfWebAcl_CognitoUserPoolProperty
	_jsii_.Get(
		j,
		"cognitoUserPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) VerifiedAccessInstance() TfWebAcl_VerifiedAccessInstancePropertyOutputReference {
	var returns TfWebAcl_VerifiedAccessInstancePropertyOutputReference
	_jsii_.Get(
		j,
		"verifiedAccessInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) VerifiedAccessInstanceInput() *TfWebAcl_VerifiedAccessInstanceProperty {
	var returns *TfWebAcl_VerifiedAccessInstanceProperty
	_jsii_.Get(
		j,
		"verifiedAccessInstanceInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAcl_RequestBodyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAcl_RequestBodyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAcl_RequestBodyPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAcl.RequestBodyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAcl_RequestBodyPropertyOutputReference_Override(t TfWebAcl_RequestBodyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAcl.RequestBodyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) PutApiGateway(value *TfWebAcl_ApiGatewayProperty) {
	if err := t.validatePutApiGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putApiGateway",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) PutAppRunnerService(value *TfWebAcl_AppRunnerServiceProperty) {
	if err := t.validatePutAppRunnerServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAppRunnerService",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) PutCloudfront(value *TfWebAcl_CloudfrontProperty) {
	if err := t.validatePutCloudfrontParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudfront",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) PutCognitoUserPool(value *TfWebAcl_CognitoUserPoolProperty) {
	if err := t.validatePutCognitoUserPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCognitoUserPool",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) PutVerifiedAccessInstance(value *TfWebAcl_VerifiedAccessInstanceProperty) {
	if err := t.validatePutVerifiedAccessInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVerifiedAccessInstance",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) ResetApiGateway() {
	_jsii_.InvokeVoid(
		t,
		"resetApiGateway",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) ResetAppRunnerService() {
	_jsii_.InvokeVoid(
		t,
		"resetAppRunnerService",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) ResetCloudfront() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudfront",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) ResetCognitoUserPool() {
	_jsii_.InvokeVoid(
		t,
		"resetCognitoUserPool",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) ResetVerifiedAccessInstance() {
	_jsii_.InvokeVoid(
		t,
		"resetVerifiedAccessInstance",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAcl_RequestBodyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

