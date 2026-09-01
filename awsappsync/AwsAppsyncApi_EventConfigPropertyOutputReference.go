package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappsync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappsync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppsyncApi_EventConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthProvider() AwsAppsyncApi_AuthProviderPropertyList
	// Experimental.
	AuthProviderInput() interface{}
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
	ConnectionAuthMode() AwsAppsyncApi_ConnectionAuthModePropertyList
	// Experimental.
	ConnectionAuthModeInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DefaultPublishAuthMode() AwsAppsyncApi_DefaultPublishAuthModePropertyList
	// Experimental.
	DefaultPublishAuthModeInput() interface{}
	// Experimental.
	DefaultSubscribeAuthMode() AwsAppsyncApi_DefaultSubscribeAuthModePropertyList
	// Experimental.
	DefaultSubscribeAuthModeInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LogConfig() AwsAppsyncApi_LogConfigPropertyList
	// Experimental.
	LogConfigInput() interface{}
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
	PutAuthProvider(value interface{})
	// Experimental.
	PutConnectionAuthMode(value interface{})
	// Experimental.
	PutDefaultPublishAuthMode(value interface{})
	// Experimental.
	PutDefaultSubscribeAuthMode(value interface{})
	// Experimental.
	PutLogConfig(value interface{})
	// Experimental.
	ResetAuthProvider()
	// Experimental.
	ResetConnectionAuthMode()
	// Experimental.
	ResetDefaultPublishAuthMode()
	// Experimental.
	ResetDefaultSubscribeAuthMode()
	// Experimental.
	ResetLogConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsAppsyncApi_EventConfigPropertyOutputReference
type jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) AuthProvider() AwsAppsyncApi_AuthProviderPropertyList {
	var returns AwsAppsyncApi_AuthProviderPropertyList
	_jsii_.Get(
		j,
		"authProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) AuthProviderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) ConnectionAuthMode() AwsAppsyncApi_ConnectionAuthModePropertyList {
	var returns AwsAppsyncApi_ConnectionAuthModePropertyList
	_jsii_.Get(
		j,
		"connectionAuthMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) ConnectionAuthModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionAuthModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) DefaultPublishAuthMode() AwsAppsyncApi_DefaultPublishAuthModePropertyList {
	var returns AwsAppsyncApi_DefaultPublishAuthModePropertyList
	_jsii_.Get(
		j,
		"defaultPublishAuthMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) DefaultPublishAuthModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultPublishAuthModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) DefaultSubscribeAuthMode() AwsAppsyncApi_DefaultSubscribeAuthModePropertyList {
	var returns AwsAppsyncApi_DefaultSubscribeAuthModePropertyList
	_jsii_.Get(
		j,
		"defaultSubscribeAuthMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) DefaultSubscribeAuthModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultSubscribeAuthModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) LogConfig() AwsAppsyncApi_LogConfigPropertyList {
	var returns AwsAppsyncApi_LogConfigPropertyList
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) LogConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppsyncApi_EventConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsAppsyncApi_EventConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppsyncApi_EventConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsAppsyncApi.EventConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppsyncApi_EventConfigPropertyOutputReference_Override(a AwsAppsyncApi_EventConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsAppsyncApi.EventConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) PutAuthProvider(value interface{}) {
	if err := a.validatePutAuthProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthProvider",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) PutConnectionAuthMode(value interface{}) {
	if err := a.validatePutConnectionAuthModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectionAuthMode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) PutDefaultPublishAuthMode(value interface{}) {
	if err := a.validatePutDefaultPublishAuthModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultPublishAuthMode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) PutDefaultSubscribeAuthMode(value interface{}) {
	if err := a.validatePutDefaultSubscribeAuthModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultSubscribeAuthMode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) PutLogConfig(value interface{}) {
	if err := a.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) ResetAuthProvider() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthProvider",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) ResetConnectionAuthMode() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionAuthMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) ResetDefaultPublishAuthMode() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultPublishAuthMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) ResetDefaultSubscribeAuthMode() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultSubscribeAuthMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) ResetLogConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppsyncApi_EventConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

