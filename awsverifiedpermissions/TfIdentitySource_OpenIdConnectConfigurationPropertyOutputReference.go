package awsverifiedpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsverifiedpermissions/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsverifiedpermissions/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference interface {
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
	EntityIdPrefix() *string
	// Experimental.
	SetEntityIdPrefix(val *string)
	// Experimental.
	EntityIdPrefixInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	GroupConfiguration() TfIdentitySource_ConfigurationOpenIdConnectConfigurationGroupConfigurationPropertyList
	// Experimental.
	GroupConfigurationInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Issuer() *string
	// Experimental.
	SetIssuer(val *string)
	// Experimental.
	IssuerInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TokenSelection() TfIdentitySource_TokenSelectionPropertyList
	// Experimental.
	TokenSelectionInput() interface{}
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
	PutTokenSelection(value interface{})
	// Experimental.
	ResetEntityIdPrefix()
	// Experimental.
	ResetGroupConfiguration()
	// Experimental.
	ResetTokenSelection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference
type jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) EntityIdPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityIdPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) EntityIdPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityIdPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GroupConfiguration() TfIdentitySource_ConfigurationOpenIdConnectConfigurationGroupConfigurationPropertyList {
	var returns TfIdentitySource_ConfigurationOpenIdConnectConfigurationGroupConfigurationPropertyList
	_jsii_.Get(
		j,
		"groupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GroupConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) TokenSelection() TfIdentitySource_TokenSelectionPropertyList {
	var returns TfIdentitySource_TokenSelectionPropertyList
	_jsii_.Get(
		j,
		"tokenSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) TokenSelectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenSelectionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfIdentitySource_OpenIdConnectConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-verified-permissions.TfIdentitySource.OpenIdConnectConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference_Override(t TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-verified-permissions.TfIdentitySource.OpenIdConnectConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference)SetEntityIdPrefix(val *string) {
	if err := j.validateSetEntityIdPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityIdPrefix",
		val,
	)
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) PutGroupConfiguration(value interface{}) {
	if err := t.validatePutGroupConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGroupConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) PutTokenSelection(value interface{}) {
	if err := t.validatePutTokenSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTokenSelection",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) ResetEntityIdPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetEntityIdPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) ResetGroupConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetGroupConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) ResetTokenSelection() {
	_jsii_.InvokeVoid(
		t,
		"resetTokenSelection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

