package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference interface {
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
	DefaultEmailOption() *string
	// Experimental.
	SetDefaultEmailOption(val *string)
	// Experimental.
	DefaultEmailOptionInput() *string
	// Experimental.
	EmailMessage() *string
	// Experimental.
	SetEmailMessage(val *string)
	// Experimental.
	EmailMessageByLink() *string
	// Experimental.
	SetEmailMessageByLink(val *string)
	// Experimental.
	EmailMessageByLinkInput() *string
	// Experimental.
	EmailMessageInput() *string
	// Experimental.
	EmailSubject() *string
	// Experimental.
	SetEmailSubject(val *string)
	// Experimental.
	EmailSubjectByLink() *string
	// Experimental.
	SetEmailSubjectByLink(val *string)
	// Experimental.
	EmailSubjectByLinkInput() *string
	// Experimental.
	EmailSubjectInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsCognitoUserPool_VerificationMessageTemplateProperty
	// Experimental.
	SetInternalValue(val *AwsCognitoUserPool_VerificationMessageTemplateProperty)
	// Experimental.
	SmsMessage() *string
	// Experimental.
	SetSmsMessage(val *string)
	// Experimental.
	SmsMessageInput() *string
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
	ResetDefaultEmailOption()
	// Experimental.
	ResetEmailMessage()
	// Experimental.
	ResetEmailMessageByLink()
	// Experimental.
	ResetEmailSubject()
	// Experimental.
	ResetEmailSubjectByLink()
	// Experimental.
	ResetSmsMessage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference
type jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) DefaultEmailOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultEmailOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) DefaultEmailOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultEmailOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) EmailMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) EmailMessageByLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessageByLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) EmailMessageByLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessageByLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) EmailMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) EmailSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) EmailSubjectByLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubjectByLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) EmailSubjectByLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubjectByLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) EmailSubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) InternalValue() *AwsCognitoUserPool_VerificationMessageTemplateProperty {
	var returns *AwsCognitoUserPool_VerificationMessageTemplateProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) SmsMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) SmsMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoUserPool.VerificationMessageTemplatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference_Override(a AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoUserPool.VerificationMessageTemplatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference)SetDefaultEmailOption(val *string) {
	if err := j.validateSetDefaultEmailOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultEmailOption",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference)SetEmailMessage(val *string) {
	if err := j.validateSetEmailMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailMessage",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference)SetEmailMessageByLink(val *string) {
	if err := j.validateSetEmailMessageByLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailMessageByLink",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference)SetEmailSubject(val *string) {
	if err := j.validateSetEmailSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailSubject",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference)SetEmailSubjectByLink(val *string) {
	if err := j.validateSetEmailSubjectByLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailSubjectByLink",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference)SetInternalValue(val *AwsCognitoUserPool_VerificationMessageTemplateProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference)SetSmsMessage(val *string) {
	if err := j.validateSetSmsMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smsMessage",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) ResetDefaultEmailOption() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultEmailOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) ResetEmailMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) ResetEmailMessageByLink() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailMessageByLink",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) ResetEmailSubject() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailSubject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) ResetEmailSubjectByLink() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailSubjectByLink",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) ResetSmsMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetSmsMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCognitoUserPool_VerificationMessageTemplatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

