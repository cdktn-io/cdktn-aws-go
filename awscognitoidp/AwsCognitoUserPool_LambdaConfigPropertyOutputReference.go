package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCognitoUserPool_LambdaConfigPropertyOutputReference interface {
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
	CustomEmailSender() AwsCognitoUserPool_CustomEmailSenderPropertyOutputReference
	// Experimental.
	CustomEmailSenderInput() *AwsCognitoUserPool_CustomEmailSenderProperty
	// Experimental.
	CustomMessage() *string
	// Experimental.
	SetCustomMessage(val *string)
	// Experimental.
	CustomMessageInput() *string
	// Experimental.
	CustomSmsSender() AwsCognitoUserPool_CustomSmsSenderPropertyOutputReference
	// Experimental.
	CustomSmsSenderInput() *AwsCognitoUserPool_CustomSmsSenderProperty
	// Experimental.
	DefineAuthChallenge() *string
	// Experimental.
	SetDefineAuthChallenge(val *string)
	// Experimental.
	DefineAuthChallengeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsCognitoUserPool_LambdaConfigProperty
	// Experimental.
	SetInternalValue(val *AwsCognitoUserPool_LambdaConfigProperty)
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
	PreTokenGenerationConfig() AwsCognitoUserPool_PreTokenGenerationConfigPropertyOutputReference
	// Experimental.
	PreTokenGenerationConfigInput() *AwsCognitoUserPool_PreTokenGenerationConfigProperty
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
	PutCustomEmailSender(value *AwsCognitoUserPool_CustomEmailSenderProperty)
	// Experimental.
	PutCustomSmsSender(value *AwsCognitoUserPool_CustomSmsSenderProperty)
	// Experimental.
	PutPreTokenGenerationConfig(value *AwsCognitoUserPool_PreTokenGenerationConfigProperty)
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

// The jsii proxy struct for AwsCognitoUserPool_LambdaConfigPropertyOutputReference
type jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) CreateAuthChallenge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAuthChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) CreateAuthChallengeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAuthChallengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) CustomEmailSender() AwsCognitoUserPool_CustomEmailSenderPropertyOutputReference {
	var returns AwsCognitoUserPool_CustomEmailSenderPropertyOutputReference
	_jsii_.Get(
		j,
		"customEmailSender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) CustomEmailSenderInput() *AwsCognitoUserPool_CustomEmailSenderProperty {
	var returns *AwsCognitoUserPool_CustomEmailSenderProperty
	_jsii_.Get(
		j,
		"customEmailSenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) CustomMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) CustomMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) CustomSmsSender() AwsCognitoUserPool_CustomSmsSenderPropertyOutputReference {
	var returns AwsCognitoUserPool_CustomSmsSenderPropertyOutputReference
	_jsii_.Get(
		j,
		"customSmsSender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) CustomSmsSenderInput() *AwsCognitoUserPool_CustomSmsSenderProperty {
	var returns *AwsCognitoUserPool_CustomSmsSenderProperty
	_jsii_.Get(
		j,
		"customSmsSenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) DefineAuthChallenge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defineAuthChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) DefineAuthChallengeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defineAuthChallengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) InternalValue() *AwsCognitoUserPool_LambdaConfigProperty {
	var returns *AwsCognitoUserPool_LambdaConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) PostAuthentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) PostAuthenticationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) PostConfirmation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postConfirmation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) PostConfirmationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postConfirmationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) PreAuthentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) PreAuthenticationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) PreSignUp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSignUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) PreSignUpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSignUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) PreTokenGeneration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preTokenGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) PreTokenGenerationConfig() AwsCognitoUserPool_PreTokenGenerationConfigPropertyOutputReference {
	var returns AwsCognitoUserPool_PreTokenGenerationConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"preTokenGenerationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) PreTokenGenerationConfigInput() *AwsCognitoUserPool_PreTokenGenerationConfigProperty {
	var returns *AwsCognitoUserPool_PreTokenGenerationConfigProperty
	_jsii_.Get(
		j,
		"preTokenGenerationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) PreTokenGenerationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preTokenGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) UserMigration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userMigration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) UserMigrationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userMigrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) VerifyAuthChallengeResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyAuthChallengeResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) VerifyAuthChallengeResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyAuthChallengeResponseInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCognitoUserPool_LambdaConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCognitoUserPool_LambdaConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCognitoUserPool_LambdaConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoUserPool.LambdaConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCognitoUserPool_LambdaConfigPropertyOutputReference_Override(a AwsCognitoUserPool_LambdaConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoUserPool.LambdaConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetCreateAuthChallenge(val *string) {
	if err := j.validateSetCreateAuthChallengeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAuthChallenge",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetCustomMessage(val *string) {
	if err := j.validateSetCustomMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customMessage",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetDefineAuthChallenge(val *string) {
	if err := j.validateSetDefineAuthChallengeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defineAuthChallenge",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetInternalValue(val *AwsCognitoUserPool_LambdaConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetPostAuthentication(val *string) {
	if err := j.validateSetPostAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postAuthentication",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetPostConfirmation(val *string) {
	if err := j.validateSetPostConfirmationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postConfirmation",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetPreAuthentication(val *string) {
	if err := j.validateSetPreAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preAuthentication",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetPreSignUp(val *string) {
	if err := j.validateSetPreSignUpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preSignUp",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetPreTokenGeneration(val *string) {
	if err := j.validateSetPreTokenGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preTokenGeneration",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetUserMigration(val *string) {
	if err := j.validateSetUserMigrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userMigration",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetVerifyAuthChallengeResponse(val *string) {
	if err := j.validateSetVerifyAuthChallengeResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifyAuthChallengeResponse",
		val,
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) PutCustomEmailSender(value *AwsCognitoUserPool_CustomEmailSenderProperty) {
	if err := a.validatePutCustomEmailSenderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomEmailSender",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) PutCustomSmsSender(value *AwsCognitoUserPool_CustomSmsSenderProperty) {
	if err := a.validatePutCustomSmsSenderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomSmsSender",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) PutPreTokenGenerationConfig(value *AwsCognitoUserPool_PreTokenGenerationConfigProperty) {
	if err := a.validatePutPreTokenGenerationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPreTokenGenerationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ResetCreateAuthChallenge() {
	_jsii_.InvokeVoid(
		a,
		"resetCreateAuthChallenge",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ResetCustomEmailSender() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomEmailSender",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ResetCustomMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ResetCustomSmsSender() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomSmsSender",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ResetDefineAuthChallenge() {
	_jsii_.InvokeVoid(
		a,
		"resetDefineAuthChallenge",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ResetPostAuthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetPostAuthentication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ResetPostConfirmation() {
	_jsii_.InvokeVoid(
		a,
		"resetPostConfirmation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ResetPreAuthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetPreAuthentication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ResetPreSignUp() {
	_jsii_.InvokeVoid(
		a,
		"resetPreSignUp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ResetPreTokenGeneration() {
	_jsii_.InvokeVoid(
		a,
		"resetPreTokenGeneration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ResetPreTokenGenerationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetPreTokenGenerationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ResetUserMigration() {
	_jsii_.InvokeVoid(
		a,
		"resetUserMigration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ResetVerifyAuthChallengeResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetVerifyAuthChallengeResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCognitoUserPool_LambdaConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

