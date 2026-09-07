package appflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference interface {
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
	InternalValue() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty
	// Experimental.
	SetInternalValue(val *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty)
	// Experimental.
	LogonLanguage() *string
	// Experimental.
	SetLogonLanguage(val *string)
	// Experimental.
	LogonLanguageInput() *string
	// Experimental.
	OauthProperties() AwsConnectorProfile_OauthPropertiesPropertyOutputReference
	// Experimental.
	OauthPropertiesInput() *AwsConnectorProfile_OauthPropertiesProperty
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
	PutOauthProperties(value *AwsConnectorProfile_OauthPropertiesProperty)
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

// The jsii proxy struct for AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference
type jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ApplicationHostUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationHostUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ApplicationHostUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationHostUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ApplicationServicePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationServicePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ApplicationServicePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationServicePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ClientNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ClientNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) InternalValue() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) LogonLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logonLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) LogonLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logonLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) OauthProperties() AwsConnectorProfile_OauthPropertiesPropertyOutputReference {
	var returns AwsConnectorProfile_OauthPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"oauthProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) OauthPropertiesInput() *AwsConnectorProfile_OauthPropertiesProperty {
	var returns *AwsConnectorProfile_OauthPropertiesProperty
	_jsii_.Get(
		j,
		"oauthPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) PortNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) PortNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) PrivateLinkServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) PrivateLinkServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkServiceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsConnectorProfile.ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference_Override(a AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsConnectorProfile.ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetApplicationHostUrl(val *string) {
	if err := j.validateSetApplicationHostUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationHostUrl",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetApplicationServicePath(val *string) {
	if err := j.validateSetApplicationServicePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationServicePath",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetClientNumber(val *string) {
	if err := j.validateSetClientNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientNumber",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetInternalValue(val *AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetLogonLanguage(val *string) {
	if err := j.validateSetLogonLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logonLanguage",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetPortNumber(val *float64) {
	if err := j.validateSetPortNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portNumber",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetPrivateLinkServiceName(val *string) {
	if err := j.validateSetPrivateLinkServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateLinkServiceName",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) PutOauthProperties(value *AwsConnectorProfile_OauthPropertiesProperty) {
	if err := a.validatePutOauthPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOauthProperties",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ResetLogonLanguage() {
	_jsii_.InvokeVoid(
		a,
		"resetLogonLanguage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ResetOauthProperties() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthProperties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ResetPrivateLinkServiceName() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateLinkServiceName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

