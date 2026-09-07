package verifiedpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/verifiedpermissions/jsii"

	"github.com/cdktn-io/cdktn-aws-go/verifiedpermissions/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference interface {
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
	GroupConfiguration() AwsIdentitySource_ConfigurationOpenIdConnectConfigurationGroupConfigurationPropertyList
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
	TokenSelection() AwsIdentitySource_TokenSelectionPropertyList
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

// The jsii proxy struct for AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference
type jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) EntityIdPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityIdPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) EntityIdPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityIdPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GroupConfiguration() AwsIdentitySource_ConfigurationOpenIdConnectConfigurationGroupConfigurationPropertyList {
	var returns AwsIdentitySource_ConfigurationOpenIdConnectConfigurationGroupConfigurationPropertyList
	_jsii_.Get(
		j,
		"groupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GroupConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) TokenSelection() AwsIdentitySource_TokenSelectionPropertyList {
	var returns AwsIdentitySource_TokenSelectionPropertyList
	_jsii_.Get(
		j,
		"tokenSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) TokenSelectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenSelectionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-verified-permissions.AwsIdentitySource.OpenIdConnectConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference_Override(a AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-verified-permissions.AwsIdentitySource.OpenIdConnectConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference)SetEntityIdPrefix(val *string) {
	if err := j.validateSetEntityIdPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityIdPrefix",
		val,
	)
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) PutGroupConfiguration(value interface{}) {
	if err := a.validatePutGroupConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGroupConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) PutTokenSelection(value interface{}) {
	if err := a.validatePutTokenSelectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTokenSelection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) ResetEntityIdPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetEntityIdPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) ResetGroupConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) ResetTokenSelection() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenSelection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsIdentitySource_OpenIdConnectConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

