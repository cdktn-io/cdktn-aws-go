package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BlockEmail() AwsCognitoRiskConfiguration_BlockEmailPropertyOutputReference
	// Experimental.
	BlockEmailInput() *AwsCognitoRiskConfiguration_BlockEmailProperty
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
	From() *string
	// Experimental.
	SetFrom(val *string)
	// Experimental.
	FromInput() *string
	// Experimental.
	InternalValue() *AwsCognitoRiskConfiguration_NotifyConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsCognitoRiskConfiguration_NotifyConfigurationProperty)
	// Experimental.
	MfaEmail() AwsCognitoRiskConfiguration_MfaEmailPropertyOutputReference
	// Experimental.
	MfaEmailInput() *AwsCognitoRiskConfiguration_MfaEmailProperty
	// Experimental.
	NoActionEmail() AwsCognitoRiskConfiguration_NoActionEmailPropertyOutputReference
	// Experimental.
	NoActionEmailInput() *AwsCognitoRiskConfiguration_NoActionEmailProperty
	// Experimental.
	ReplyTo() *string
	// Experimental.
	SetReplyTo(val *string)
	// Experimental.
	ReplyToInput() *string
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
	PutBlockEmail(value *AwsCognitoRiskConfiguration_BlockEmailProperty)
	// Experimental.
	PutMfaEmail(value *AwsCognitoRiskConfiguration_MfaEmailProperty)
	// Experimental.
	PutNoActionEmail(value *AwsCognitoRiskConfiguration_NoActionEmailProperty)
	// Experimental.
	ResetBlockEmail()
	// Experimental.
	ResetFrom()
	// Experimental.
	ResetMfaEmail()
	// Experimental.
	ResetNoActionEmail()
	// Experimental.
	ResetReplyTo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference
type jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) BlockEmail() AwsCognitoRiskConfiguration_BlockEmailPropertyOutputReference {
	var returns AwsCognitoRiskConfiguration_BlockEmailPropertyOutputReference
	_jsii_.Get(
		j,
		"blockEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) BlockEmailInput() *AwsCognitoRiskConfiguration_BlockEmailProperty {
	var returns *AwsCognitoRiskConfiguration_BlockEmailProperty
	_jsii_.Get(
		j,
		"blockEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) From() *string {
	var returns *string
	_jsii_.Get(
		j,
		"from",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) FromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) InternalValue() *AwsCognitoRiskConfiguration_NotifyConfigurationProperty {
	var returns *AwsCognitoRiskConfiguration_NotifyConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) MfaEmail() AwsCognitoRiskConfiguration_MfaEmailPropertyOutputReference {
	var returns AwsCognitoRiskConfiguration_MfaEmailPropertyOutputReference
	_jsii_.Get(
		j,
		"mfaEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) MfaEmailInput() *AwsCognitoRiskConfiguration_MfaEmailProperty {
	var returns *AwsCognitoRiskConfiguration_MfaEmailProperty
	_jsii_.Get(
		j,
		"mfaEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) NoActionEmail() AwsCognitoRiskConfiguration_NoActionEmailPropertyOutputReference {
	var returns AwsCognitoRiskConfiguration_NoActionEmailPropertyOutputReference
	_jsii_.Get(
		j,
		"noActionEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) NoActionEmailInput() *AwsCognitoRiskConfiguration_NoActionEmailProperty {
	var returns *AwsCognitoRiskConfiguration_NoActionEmailProperty
	_jsii_.Get(
		j,
		"noActionEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) ReplyTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) ReplyToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) SourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoRiskConfiguration.NotifyConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference_Override(a AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoRiskConfiguration.NotifyConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetFrom(val *string) {
	if err := j.validateSetFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"from",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetInternalValue(val *AwsCognitoRiskConfiguration_NotifyConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetReplyTo(val *string) {
	if err := j.validateSetReplyToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replyTo",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetSourceArn(val *string) {
	if err := j.validateSetSourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceArn",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) PutBlockEmail(value *AwsCognitoRiskConfiguration_BlockEmailProperty) {
	if err := a.validatePutBlockEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBlockEmail",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) PutMfaEmail(value *AwsCognitoRiskConfiguration_MfaEmailProperty) {
	if err := a.validatePutMfaEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMfaEmail",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) PutNoActionEmail(value *AwsCognitoRiskConfiguration_NoActionEmailProperty) {
	if err := a.validatePutNoActionEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNoActionEmail",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) ResetBlockEmail() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockEmail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) ResetFrom() {
	_jsii_.InvokeVoid(
		a,
		"resetFrom",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) ResetMfaEmail() {
	_jsii_.InvokeVoid(
		a,
		"resetMfaEmail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) ResetNoActionEmail() {
	_jsii_.InvokeVoid(
		a,
		"resetNoActionEmail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) ResetReplyTo() {
	_jsii_.InvokeVoid(
		a,
		"resetReplyTo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCognitoRiskConfiguration_NotifyConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

