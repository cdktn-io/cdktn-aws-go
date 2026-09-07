package cognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BlockEmail() AwsRiskConfiguration_BlockEmailPropertyOutputReference
	// Experimental.
	BlockEmailInput() *AwsRiskConfiguration_BlockEmailProperty
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
	InternalValue() *AwsRiskConfiguration_NotifyConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsRiskConfiguration_NotifyConfigurationProperty)
	// Experimental.
	MfaEmail() AwsRiskConfiguration_MfaEmailPropertyOutputReference
	// Experimental.
	MfaEmailInput() *AwsRiskConfiguration_MfaEmailProperty
	// Experimental.
	NoActionEmail() AwsRiskConfiguration_NoActionEmailPropertyOutputReference
	// Experimental.
	NoActionEmailInput() *AwsRiskConfiguration_NoActionEmailProperty
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
	PutBlockEmail(value *AwsRiskConfiguration_BlockEmailProperty)
	// Experimental.
	PutMfaEmail(value *AwsRiskConfiguration_MfaEmailProperty)
	// Experimental.
	PutNoActionEmail(value *AwsRiskConfiguration_NoActionEmailProperty)
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

// The jsii proxy struct for AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference
type jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) BlockEmail() AwsRiskConfiguration_BlockEmailPropertyOutputReference {
	var returns AwsRiskConfiguration_BlockEmailPropertyOutputReference
	_jsii_.Get(
		j,
		"blockEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) BlockEmailInput() *AwsRiskConfiguration_BlockEmailProperty {
	var returns *AwsRiskConfiguration_BlockEmailProperty
	_jsii_.Get(
		j,
		"blockEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) From() *string {
	var returns *string
	_jsii_.Get(
		j,
		"from",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) FromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) InternalValue() *AwsRiskConfiguration_NotifyConfigurationProperty {
	var returns *AwsRiskConfiguration_NotifyConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) MfaEmail() AwsRiskConfiguration_MfaEmailPropertyOutputReference {
	var returns AwsRiskConfiguration_MfaEmailPropertyOutputReference
	_jsii_.Get(
		j,
		"mfaEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) MfaEmailInput() *AwsRiskConfiguration_MfaEmailProperty {
	var returns *AwsRiskConfiguration_MfaEmailProperty
	_jsii_.Get(
		j,
		"mfaEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) NoActionEmail() AwsRiskConfiguration_NoActionEmailPropertyOutputReference {
	var returns AwsRiskConfiguration_NoActionEmailPropertyOutputReference
	_jsii_.Get(
		j,
		"noActionEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) NoActionEmailInput() *AwsRiskConfiguration_NoActionEmailProperty {
	var returns *AwsRiskConfiguration_NoActionEmailProperty
	_jsii_.Get(
		j,
		"noActionEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) ReplyTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) ReplyToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) SourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRiskConfiguration_NotifyConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRiskConfiguration_NotifyConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsRiskConfiguration.NotifyConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRiskConfiguration_NotifyConfigurationPropertyOutputReference_Override(a AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsRiskConfiguration.NotifyConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetFrom(val *string) {
	if err := j.validateSetFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"from",
		val,
	)
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetInternalValue(val *AwsRiskConfiguration_NotifyConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetReplyTo(val *string) {
	if err := j.validateSetReplyToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replyTo",
		val,
	)
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetSourceArn(val *string) {
	if err := j.validateSetSourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceArn",
		val,
	)
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) PutBlockEmail(value *AwsRiskConfiguration_BlockEmailProperty) {
	if err := a.validatePutBlockEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBlockEmail",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) PutMfaEmail(value *AwsRiskConfiguration_MfaEmailProperty) {
	if err := a.validatePutMfaEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMfaEmail",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) PutNoActionEmail(value *AwsRiskConfiguration_NoActionEmailProperty) {
	if err := a.validatePutNoActionEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNoActionEmail",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) ResetBlockEmail() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockEmail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) ResetFrom() {
	_jsii_.InvokeVoid(
		a,
		"resetFrom",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) ResetMfaEmail() {
	_jsii_.InvokeVoid(
		a,
		"resetMfaEmail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) ResetNoActionEmail() {
	_jsii_.InvokeVoid(
		a,
		"resetNoActionEmail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) ResetReplyTo() {
	_jsii_.InvokeVoid(
		a,
		"resetReplyTo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRiskConfiguration_NotifyConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

