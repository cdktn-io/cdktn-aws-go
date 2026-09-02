package awsverifiedpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsverifiedpermissions/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsverifiedpermissions/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ClientIds() *[]*string
	// Experimental.
	SetClientIds(val *[]*string)
	// Experimental.
	ClientIdsInput() *[]*string
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
	GroupConfiguration() TfIdentitySource_ConfigurationCognitoUserPoolConfigurationGroupConfigurationPropertyList
	// Experimental.
	GroupConfigurationInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UserPoolArn() *string
	// Experimental.
	SetUserPoolArn(val *string)
	// Experimental.
	UserPoolArnInput() *string
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
	PutGroupConfiguration(value interface{})
	// Experimental.
	ResetClientIds()
	// Experimental.
	ResetGroupConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference
type jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) ClientIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) ClientIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) GroupConfiguration() TfIdentitySource_ConfigurationCognitoUserPoolConfigurationGroupConfigurationPropertyList {
	var returns TfIdentitySource_ConfigurationCognitoUserPoolConfigurationGroupConfigurationPropertyList
	_jsii_.Get(
		j,
		"groupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) GroupConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) UserPoolArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) UserPoolArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolArnInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-verified-permissions.TfIdentitySource.CognitoUserPoolConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference_Override(t TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-verified-permissions.TfIdentitySource.CognitoUserPoolConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference)SetClientIds(val *[]*string) {
	if err := j.validateSetClientIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientIds",
		val,
	)
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference)SetUserPoolArn(val *string) {
	if err := j.validateSetUserPoolArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolArn",
		val,
	)
}

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) PutGroupConfiguration(value interface{}) {
	if err := t.validatePutGroupConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGroupConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) ResetClientIds() {
	_jsii_.InvokeVoid(
		t,
		"resetClientIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) ResetGroupConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetGroupConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfIdentitySource_CognitoUserPoolConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

