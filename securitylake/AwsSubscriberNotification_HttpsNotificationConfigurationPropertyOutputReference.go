package securitylake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/securitylake/jsii"

	"github.com/cdktn-io/cdktn-aws-go/securitylake/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthorizationApiKeyName() *string
	// Experimental.
	SetAuthorizationApiKeyName(val *string)
	// Experimental.
	AuthorizationApiKeyNameInput() *string
	// Experimental.
	AuthorizationApiKeyValue() *string
	// Experimental.
	SetAuthorizationApiKeyValue(val *string)
	// Experimental.
	AuthorizationApiKeyValueInput() *string
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
	Endpoint() *string
	// Experimental.
	SetEndpoint(val *string)
	// Experimental.
	EndpointInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	HttpMethod() *string
	// Experimental.
	SetHttpMethod(val *string)
	// Experimental.
	HttpMethodInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	TargetRoleArn() *string
	// Experimental.
	SetTargetRoleArn(val *string)
	// Experimental.
	TargetRoleArnInput() *string
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
	ResetAuthorizationApiKeyName()
	// Experimental.
	ResetAuthorizationApiKeyValue()
	// Experimental.
	ResetHttpMethod()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference
type jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) AuthorizationApiKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationApiKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) AuthorizationApiKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationApiKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) AuthorizationApiKeyValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationApiKeyValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) AuthorizationApiKeyValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationApiKeyValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) TargetRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) TargetRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-security-lake.AwsSubscriberNotification.HttpsNotificationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference_Override(a AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-security-lake.AwsSubscriberNotification.HttpsNotificationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference)SetAuthorizationApiKeyName(val *string) {
	if err := j.validateSetAuthorizationApiKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationApiKeyName",
		val,
	)
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference)SetAuthorizationApiKeyValue(val *string) {
	if err := j.validateSetAuthorizationApiKeyValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationApiKeyValue",
		val,
	)
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference)SetEndpoint(val *string) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference)SetTargetRoleArn(val *string) {
	if err := j.validateSetTargetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) ResetAuthorizationApiKeyName() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizationApiKeyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) ResetAuthorizationApiKeyValue() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizationApiKeyValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSubscriberNotification_HttpsNotificationConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

