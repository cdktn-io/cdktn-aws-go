package cognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsUserPool_EmailConfigurationPropertyOutputReference interface {
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
	ConfigurationSet() *string
	// Experimental.
	SetConfigurationSet(val *string)
	// Experimental.
	ConfigurationSetInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EmailSendingAccount() *string
	// Experimental.
	SetEmailSendingAccount(val *string)
	// Experimental.
	EmailSendingAccountInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FromEmailAddress() *string
	// Experimental.
	SetFromEmailAddress(val *string)
	// Experimental.
	FromEmailAddressInput() *string
	// Experimental.
	InternalValue() *AwsUserPool_EmailConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsUserPool_EmailConfigurationProperty)
	// Experimental.
	ReplyToEmailAddress() *string
	// Experimental.
	SetReplyToEmailAddress(val *string)
	// Experimental.
	ReplyToEmailAddressInput() *string
	// Experimental.
	SourceArn() *string
	// Experimental.
	SetSourceArn(val *string)
	// Experimental.
	SourceArnInput() *string
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
	ResetConfigurationSet()
	// Experimental.
	ResetEmailSendingAccount()
	// Experimental.
	ResetFromEmailAddress()
	// Experimental.
	ResetReplyToEmailAddress()
	// Experimental.
	ResetSourceArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsUserPool_EmailConfigurationPropertyOutputReference
type jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) ConfigurationSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) ConfigurationSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) EmailSendingAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSendingAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) EmailSendingAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSendingAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) FromEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) FromEmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromEmailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) InternalValue() *AwsUserPool_EmailConfigurationProperty {
	var returns *AwsUserPool_EmailConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) ReplyToEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) ReplyToEmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToEmailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) SourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsUserPool_EmailConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsUserPool_EmailConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsUserPool_EmailConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsUserPool.EmailConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsUserPool_EmailConfigurationPropertyOutputReference_Override(a AwsUserPool_EmailConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsUserPool.EmailConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference)SetConfigurationSet(val *string) {
	if err := j.validateSetConfigurationSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationSet",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference)SetEmailSendingAccount(val *string) {
	if err := j.validateSetEmailSendingAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailSendingAccount",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference)SetFromEmailAddress(val *string) {
	if err := j.validateSetFromEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fromEmailAddress",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference)SetInternalValue(val *AwsUserPool_EmailConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference)SetReplyToEmailAddress(val *string) {
	if err := j.validateSetReplyToEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replyToEmailAddress",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference)SetSourceArn(val *string) {
	if err := j.validateSetSourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceArn",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) ResetConfigurationSet() {
	_jsii_.InvokeVoid(
		a,
		"resetConfigurationSet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) ResetEmailSendingAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailSendingAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) ResetFromEmailAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetFromEmailAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) ResetReplyToEmailAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetReplyToEmailAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) ResetSourceArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsUserPool_EmailConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

