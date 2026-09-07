package connect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/connect/jsii"

	"github.com/cdktn-io/cdktn-aws-go/connect/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsQuickConnect_QuickConnectConfigPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() *AwsQuickConnect_QuickConnectConfigProperty
	// Experimental.
	SetInternalValue(val *AwsQuickConnect_QuickConnectConfigProperty)
	// Experimental.
	PhoneConfig() AwsQuickConnect_PhoneConfigPropertyList
	// Experimental.
	PhoneConfigInput() interface{}
	// Experimental.
	QueueConfig() AwsQuickConnect_QueueConfigPropertyList
	// Experimental.
	QueueConfigInput() interface{}
	// Experimental.
	QuickConnectType() *string
	// Experimental.
	SetQuickConnectType(val *string)
	// Experimental.
	QuickConnectTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UserConfig() AwsQuickConnect_UserConfigPropertyList
	// Experimental.
	UserConfigInput() interface{}
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
	PutPhoneConfig(value interface{})
	// Experimental.
	PutQueueConfig(value interface{})
	// Experimental.
	PutUserConfig(value interface{})
	// Experimental.
	ResetPhoneConfig()
	// Experimental.
	ResetQueueConfig()
	// Experimental.
	ResetUserConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsQuickConnect_QuickConnectConfigPropertyOutputReference
type jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) InternalValue() *AwsQuickConnect_QuickConnectConfigProperty {
	var returns *AwsQuickConnect_QuickConnectConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) PhoneConfig() AwsQuickConnect_PhoneConfigPropertyList {
	var returns AwsQuickConnect_PhoneConfigPropertyList
	_jsii_.Get(
		j,
		"phoneConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) PhoneConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"phoneConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) QueueConfig() AwsQuickConnect_QueueConfigPropertyList {
	var returns AwsQuickConnect_QueueConfigPropertyList
	_jsii_.Get(
		j,
		"queueConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) QueueConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queueConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) QuickConnectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quickConnectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) QuickConnectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quickConnectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) UserConfig() AwsQuickConnect_UserConfigPropertyList {
	var returns AwsQuickConnect_UserConfigPropertyList
	_jsii_.Get(
		j,
		"userConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) UserConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userConfigInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsQuickConnect_QuickConnectConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsQuickConnect_QuickConnectConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsQuickConnect_QuickConnectConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-connect.AwsQuickConnect.QuickConnectConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsQuickConnect_QuickConnectConfigPropertyOutputReference_Override(a AwsQuickConnect_QuickConnectConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect.AwsQuickConnect.QuickConnectConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference)SetInternalValue(val *AwsQuickConnect_QuickConnectConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference)SetQuickConnectType(val *string) {
	if err := j.validateSetQuickConnectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quickConnectType",
		val,
	)
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) PutPhoneConfig(value interface{}) {
	if err := a.validatePutPhoneConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPhoneConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) PutQueueConfig(value interface{}) {
	if err := a.validatePutQueueConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueueConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) PutUserConfig(value interface{}) {
	if err := a.validatePutUserConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) ResetPhoneConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetPhoneConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) ResetQueueConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetQueueConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) ResetUserConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetUserConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsQuickConnect_QuickConnectConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

