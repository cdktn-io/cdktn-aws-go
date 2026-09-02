package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApplicationHostUrl() *string
	// Experimental.
	SetApplicationHostUrl(val *string)
	// Experimental.
	ApplicationHostUrlInput() *string
	// Experimental.
	ApplicationServicePath() *string
	// Experimental.
	SetApplicationServicePath(val *string)
	// Experimental.
	ApplicationServicePathInput() *string
	// Experimental.
	ClientNumber() *string
	// Experimental.
	SetClientNumber(val *string)
	// Experimental.
	ClientNumberInput() *string
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
	InternalValue() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty
	// Experimental.
	SetInternalValue(val *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty)
	// Experimental.
	LogonLanguage() *string
	// Experimental.
	SetLogonLanguage(val *string)
	// Experimental.
	LogonLanguageInput() *string
	// Experimental.
	OauthProperties() TfConnectorProfile_OauthPropertiesPropertyOutputReference
	// Experimental.
	OauthPropertiesInput() *TfConnectorProfile_OauthPropertiesProperty
	// Experimental.
	PortNumber() *float64
	// Experimental.
	SetPortNumber(val *float64)
	// Experimental.
	PortNumberInput() *float64
	// Experimental.
	PrivateLinkServiceName() *string
	// Experimental.
	SetPrivateLinkServiceName(val *string)
	// Experimental.
	PrivateLinkServiceNameInput() *string
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
	PutOauthProperties(value *TfConnectorProfile_OauthPropertiesProperty)
	// Experimental.
	ResetLogonLanguage()
	// Experimental.
	ResetOauthProperties()
	// Experimental.
	ResetPrivateLinkServiceName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference
type jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ApplicationHostUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationHostUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ApplicationHostUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationHostUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ApplicationServicePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationServicePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ApplicationServicePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationServicePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ClientNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ClientNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) InternalValue() *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) LogonLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logonLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) LogonLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logonLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) OauthProperties() TfConnectorProfile_OauthPropertiesPropertyOutputReference {
	var returns TfConnectorProfile_OauthPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"oauthProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) OauthPropertiesInput() *TfConnectorProfile_OauthPropertiesProperty {
	var returns *TfConnectorProfile_OauthPropertiesProperty
	_jsii_.Get(
		j,
		"oauthPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) PortNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) PortNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) PrivateLinkServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) PrivateLinkServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkServiceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.TfConnectorProfile.ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference_Override(t TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.TfConnectorProfile.ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetApplicationHostUrl(val *string) {
	if err := j.validateSetApplicationHostUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationHostUrl",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetApplicationServicePath(val *string) {
	if err := j.validateSetApplicationServicePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationServicePath",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetClientNumber(val *string) {
	if err := j.validateSetClientNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientNumber",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetInternalValue(val *TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetLogonLanguage(val *string) {
	if err := j.validateSetLogonLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logonLanguage",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetPortNumber(val *float64) {
	if err := j.validateSetPortNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portNumber",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetPrivateLinkServiceName(val *string) {
	if err := j.validateSetPrivateLinkServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateLinkServiceName",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) PutOauthProperties(value *TfConnectorProfile_OauthPropertiesProperty) {
	if err := t.validatePutOauthPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOauthProperties",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ResetLogonLanguage() {
	_jsii_.InvokeVoid(
		t,
		"resetLogonLanguage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ResetOauthProperties() {
	_jsii_.InvokeVoid(
		t,
		"resetOauthProperties",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ResetPrivateLinkServiceName() {
	_jsii_.InvokeVoid(
		t,
		"resetPrivateLinkServiceName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

