package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CustomEmailSender() DataAwsCognitoUserPool_CustomEmailSenderPropertyList
	// Experimental.
	CustomMessage() *string
	// Experimental.
	CustomSmsSender() DataAwsCognitoUserPool_CustomSmsSenderPropertyList
	// Experimental.
	DefineAuthChallenge() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *DataAwsCognitoUserPool_LambdaConfigProperty
	// Experimental.
	SetInternalValue(val *DataAwsCognitoUserPool_LambdaConfigProperty)
	// Experimental.
	KmsKeyId() *string
	// Experimental.
	PostAuthentication() *string
	// Experimental.
	PostConfirmation() *string
	// Experimental.
	PreAuthentication() *string
	// Experimental.
	PreSignUp() *string
	// Experimental.
	PreTokenGeneration() *string
	// Experimental.
	PreTokenGenerationConfig() DataAwsCognitoUserPool_PreTokenGenerationConfigPropertyList
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
	VerifyAuthChallengeResponse() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference
type jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) CreateAuthChallenge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAuthChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) CustomEmailSender() DataAwsCognitoUserPool_CustomEmailSenderPropertyList {
	var returns DataAwsCognitoUserPool_CustomEmailSenderPropertyList
	_jsii_.Get(
		j,
		"customEmailSender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) CustomMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) CustomSmsSender() DataAwsCognitoUserPool_CustomSmsSenderPropertyList {
	var returns DataAwsCognitoUserPool_CustomSmsSenderPropertyList
	_jsii_.Get(
		j,
		"customSmsSender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) DefineAuthChallenge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defineAuthChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) InternalValue() *DataAwsCognitoUserPool_LambdaConfigProperty {
	var returns *DataAwsCognitoUserPool_LambdaConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) PostAuthentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) PostConfirmation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postConfirmation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) PreAuthentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) PreSignUp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSignUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) PreTokenGeneration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preTokenGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) PreTokenGenerationConfig() DataAwsCognitoUserPool_PreTokenGenerationConfigPropertyList {
	var returns DataAwsCognitoUserPool_PreTokenGenerationConfigPropertyList
	_jsii_.Get(
		j,
		"preTokenGenerationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) UserMigration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userMigration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) VerifyAuthChallengeResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyAuthChallengeResponse",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsCognitoUserPool_LambdaConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsCognitoUserPool_LambdaConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.DataAwsCognitoUserPool.LambdaConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCognitoUserPool_LambdaConfigPropertyOutputReference_Override(d DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.DataAwsCognitoUserPool.LambdaConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetInternalValue(val *DataAwsCognitoUserPool_LambdaConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCognitoUserPool_LambdaConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

