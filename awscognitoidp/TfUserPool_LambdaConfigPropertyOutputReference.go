package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserPool_LambdaConfigPropertyOutputReference interface {
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
	CreateAuthChallenge() *string
	// Experimental.
	SetCreateAuthChallenge(val *string)
	// Experimental.
	CreateAuthChallengeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CustomEmailSender() TfUserPool_CustomEmailSenderPropertyOutputReference
	// Experimental.
	CustomEmailSenderInput() *TfUserPool_CustomEmailSenderProperty
	// Experimental.
	CustomMessage() *string
	// Experimental.
	SetCustomMessage(val *string)
	// Experimental.
	CustomMessageInput() *string
	// Experimental.
	CustomSmsSender() TfUserPool_CustomSmsSenderPropertyOutputReference
	// Experimental.
	CustomSmsSenderInput() *TfUserPool_CustomSmsSenderProperty
	// Experimental.
	DefineAuthChallenge() *string
	// Experimental.
	SetDefineAuthChallenge(val *string)
	// Experimental.
	DefineAuthChallengeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfUserPool_LambdaConfigProperty
	// Experimental.
	SetInternalValue(val *TfUserPool_LambdaConfigProperty)
	// Experimental.
	KmsKeyId() *string
	// Experimental.
	SetKmsKeyId(val *string)
	// Experimental.
	KmsKeyIdInput() *string
	// Experimental.
	PostAuthentication() *string
	// Experimental.
	SetPostAuthentication(val *string)
	// Experimental.
	PostAuthenticationInput() *string
	// Experimental.
	PostConfirmation() *string
	// Experimental.
	SetPostConfirmation(val *string)
	// Experimental.
	PostConfirmationInput() *string
	// Experimental.
	PreAuthentication() *string
	// Experimental.
	SetPreAuthentication(val *string)
	// Experimental.
	PreAuthenticationInput() *string
	// Experimental.
	PreSignUp() *string
	// Experimental.
	SetPreSignUp(val *string)
	// Experimental.
	PreSignUpInput() *string
	// Experimental.
	PreTokenGeneration() *string
	// Experimental.
	SetPreTokenGeneration(val *string)
	// Experimental.
	PreTokenGenerationConfig() TfUserPool_PreTokenGenerationConfigPropertyOutputReference
	// Experimental.
	PreTokenGenerationConfigInput() *TfUserPool_PreTokenGenerationConfigProperty
	// Experimental.
	PreTokenGenerationInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UserMigration() *string
	// Experimental.
	SetUserMigration(val *string)
	// Experimental.
	UserMigrationInput() *string
	// Experimental.
	VerifyAuthChallengeResponse() *string
	// Experimental.
	SetVerifyAuthChallengeResponse(val *string)
	// Experimental.
	VerifyAuthChallengeResponseInput() *string
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
	PutCustomEmailSender(value *TfUserPool_CustomEmailSenderProperty)
	// Experimental.
	PutCustomSmsSender(value *TfUserPool_CustomSmsSenderProperty)
	// Experimental.
	PutPreTokenGenerationConfig(value *TfUserPool_PreTokenGenerationConfigProperty)
	// Experimental.
	ResetCreateAuthChallenge()
	// Experimental.
	ResetCustomEmailSender()
	// Experimental.
	ResetCustomMessage()
	// Experimental.
	ResetCustomSmsSender()
	// Experimental.
	ResetDefineAuthChallenge()
	// Experimental.
	ResetKmsKeyId()
	// Experimental.
	ResetPostAuthentication()
	// Experimental.
	ResetPostConfirmation()
	// Experimental.
	ResetPreAuthentication()
	// Experimental.
	ResetPreSignUp()
	// Experimental.
	ResetPreTokenGeneration()
	// Experimental.
	ResetPreTokenGenerationConfig()
	// Experimental.
	ResetUserMigration()
	// Experimental.
	ResetVerifyAuthChallengeResponse()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfUserPool_LambdaConfigPropertyOutputReference
type jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) CreateAuthChallenge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAuthChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) CreateAuthChallengeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAuthChallengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) CustomEmailSender() TfUserPool_CustomEmailSenderPropertyOutputReference {
	var returns TfUserPool_CustomEmailSenderPropertyOutputReference
	_jsii_.Get(
		j,
		"customEmailSender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) CustomEmailSenderInput() *TfUserPool_CustomEmailSenderProperty {
	var returns *TfUserPool_CustomEmailSenderProperty
	_jsii_.Get(
		j,
		"customEmailSenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) CustomMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) CustomMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) CustomSmsSender() TfUserPool_CustomSmsSenderPropertyOutputReference {
	var returns TfUserPool_CustomSmsSenderPropertyOutputReference
	_jsii_.Get(
		j,
		"customSmsSender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) CustomSmsSenderInput() *TfUserPool_CustomSmsSenderProperty {
	var returns *TfUserPool_CustomSmsSenderProperty
	_jsii_.Get(
		j,
		"customSmsSenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) DefineAuthChallenge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defineAuthChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) DefineAuthChallengeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defineAuthChallengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) InternalValue() *TfUserPool_LambdaConfigProperty {
	var returns *TfUserPool_LambdaConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) PostAuthentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) PostAuthenticationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) PostConfirmation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postConfirmation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) PostConfirmationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postConfirmationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) PreAuthentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) PreAuthenticationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) PreSignUp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSignUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) PreSignUpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSignUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) PreTokenGeneration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preTokenGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) PreTokenGenerationConfig() TfUserPool_PreTokenGenerationConfigPropertyOutputReference {
	var returns TfUserPool_PreTokenGenerationConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"preTokenGenerationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) PreTokenGenerationConfigInput() *TfUserPool_PreTokenGenerationConfigProperty {
	var returns *TfUserPool_PreTokenGenerationConfigProperty
	_jsii_.Get(
		j,
		"preTokenGenerationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) PreTokenGenerationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preTokenGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) UserMigration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userMigration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) UserMigrationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userMigrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) VerifyAuthChallengeResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyAuthChallengeResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) VerifyAuthChallengeResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyAuthChallengeResponseInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfUserPool_LambdaConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfUserPool_LambdaConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfUserPool_LambdaConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfUserPool.LambdaConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfUserPool_LambdaConfigPropertyOutputReference_Override(t TfUserPool_LambdaConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfUserPool.LambdaConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference)SetCreateAuthChallenge(val *string) {
	if err := j.validateSetCreateAuthChallengeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAuthChallenge",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference)SetCustomMessage(val *string) {
	if err := j.validateSetCustomMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customMessage",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference)SetDefineAuthChallenge(val *string) {
	if err := j.validateSetDefineAuthChallengeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defineAuthChallenge",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference)SetInternalValue(val *TfUserPool_LambdaConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference)SetPostAuthentication(val *string) {
	if err := j.validateSetPostAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postAuthentication",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference)SetPostConfirmation(val *string) {
	if err := j.validateSetPostConfirmationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postConfirmation",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference)SetPreAuthentication(val *string) {
	if err := j.validateSetPreAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preAuthentication",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference)SetPreSignUp(val *string) {
	if err := j.validateSetPreSignUpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preSignUp",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference)SetPreTokenGeneration(val *string) {
	if err := j.validateSetPreTokenGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preTokenGeneration",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference)SetUserMigration(val *string) {
	if err := j.validateSetUserMigrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userMigration",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference)SetVerifyAuthChallengeResponse(val *string) {
	if err := j.validateSetVerifyAuthChallengeResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifyAuthChallengeResponse",
		val,
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) PutCustomEmailSender(value *TfUserPool_CustomEmailSenderProperty) {
	if err := t.validatePutCustomEmailSenderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomEmailSender",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) PutCustomSmsSender(value *TfUserPool_CustomSmsSenderProperty) {
	if err := t.validatePutCustomSmsSenderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomSmsSender",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) PutPreTokenGenerationConfig(value *TfUserPool_PreTokenGenerationConfigProperty) {
	if err := t.validatePutPreTokenGenerationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPreTokenGenerationConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ResetCreateAuthChallenge() {
	_jsii_.InvokeVoid(
		t,
		"resetCreateAuthChallenge",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ResetCustomEmailSender() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomEmailSender",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ResetCustomMessage() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomMessage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ResetCustomSmsSender() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomSmsSender",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ResetDefineAuthChallenge() {
	_jsii_.InvokeVoid(
		t,
		"resetDefineAuthChallenge",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ResetPostAuthentication() {
	_jsii_.InvokeVoid(
		t,
		"resetPostAuthentication",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ResetPostConfirmation() {
	_jsii_.InvokeVoid(
		t,
		"resetPostConfirmation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ResetPreAuthentication() {
	_jsii_.InvokeVoid(
		t,
		"resetPreAuthentication",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ResetPreSignUp() {
	_jsii_.InvokeVoid(
		t,
		"resetPreSignUp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ResetPreTokenGeneration() {
	_jsii_.InvokeVoid(
		t,
		"resetPreTokenGeneration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ResetPreTokenGenerationConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetPreTokenGenerationConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ResetUserMigration() {
	_jsii_.InvokeVoid(
		t,
		"resetUserMigration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ResetVerifyAuthChallengeResponse() {
	_jsii_.InvokeVoid(
		t,
		"resetVerifyAuthChallengeResponse",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfUserPool_LambdaConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

