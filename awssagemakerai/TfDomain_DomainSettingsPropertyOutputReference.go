package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_DomainSettingsPropertyOutputReference interface {
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
	DockerSettings() TfDomain_DockerSettingsPropertyOutputReference
	// Experimental.
	DockerSettingsInput() *TfDomain_DockerSettingsProperty
	// Experimental.
	ExecutionRoleIdentityConfig() *string
	// Experimental.
	SetExecutionRoleIdentityConfig(val *string)
	// Experimental.
	ExecutionRoleIdentityConfigInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDomain_DomainSettingsProperty
	// Experimental.
	SetInternalValue(val *TfDomain_DomainSettingsProperty)
	// Experimental.
	RStudioServerProDomainSettings() TfDomain_RStudioServerProDomainSettingsPropertyOutputReference
	// Experimental.
	RStudioServerProDomainSettingsInput() *TfDomain_RStudioServerProDomainSettingsProperty
	// Experimental.
	SecurityGroupIds() *[]*string
	// Experimental.
	SetSecurityGroupIds(val *[]*string)
	// Experimental.
	SecurityGroupIdsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TrustedIdentityPropagationSettings() TfDomain_TrustedIdentityPropagationSettingsPropertyOutputReference
	// Experimental.
	TrustedIdentityPropagationSettingsInput() *TfDomain_TrustedIdentityPropagationSettingsProperty
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
	PutDockerSettings(value *TfDomain_DockerSettingsProperty)
	// Experimental.
	PutRStudioServerProDomainSettings(value *TfDomain_RStudioServerProDomainSettingsProperty)
	// Experimental.
	PutTrustedIdentityPropagationSettings(value *TfDomain_TrustedIdentityPropagationSettingsProperty)
	// Experimental.
	ResetDockerSettings()
	// Experimental.
	ResetExecutionRoleIdentityConfig()
	// Experimental.
	ResetRStudioServerProDomainSettings()
	// Experimental.
	ResetSecurityGroupIds()
	// Experimental.
	ResetTrustedIdentityPropagationSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDomain_DomainSettingsPropertyOutputReference
type jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) DockerSettings() TfDomain_DockerSettingsPropertyOutputReference {
	var returns TfDomain_DockerSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"dockerSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) DockerSettingsInput() *TfDomain_DockerSettingsProperty {
	var returns *TfDomain_DockerSettingsProperty
	_jsii_.Get(
		j,
		"dockerSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) ExecutionRoleIdentityConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleIdentityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) ExecutionRoleIdentityConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleIdentityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) InternalValue() *TfDomain_DomainSettingsProperty {
	var returns *TfDomain_DomainSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) RStudioServerProDomainSettings() TfDomain_RStudioServerProDomainSettingsPropertyOutputReference {
	var returns TfDomain_RStudioServerProDomainSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rStudioServerProDomainSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) RStudioServerProDomainSettingsInput() *TfDomain_RStudioServerProDomainSettingsProperty {
	var returns *TfDomain_RStudioServerProDomainSettingsProperty
	_jsii_.Get(
		j,
		"rStudioServerProDomainSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) TrustedIdentityPropagationSettings() TfDomain_TrustedIdentityPropagationSettingsPropertyOutputReference {
	var returns TfDomain_TrustedIdentityPropagationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"trustedIdentityPropagationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) TrustedIdentityPropagationSettingsInput() *TfDomain_TrustedIdentityPropagationSettingsProperty {
	var returns *TfDomain_TrustedIdentityPropagationSettingsProperty
	_jsii_.Get(
		j,
		"trustedIdentityPropagationSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_DomainSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDomain_DomainSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_DomainSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.DomainSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_DomainSettingsPropertyOutputReference_Override(t TfDomain_DomainSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.DomainSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference)SetExecutionRoleIdentityConfig(val *string) {
	if err := j.validateSetExecutionRoleIdentityConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleIdentityConfig",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference)SetInternalValue(val *TfDomain_DomainSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) PutDockerSettings(value *TfDomain_DockerSettingsProperty) {
	if err := t.validatePutDockerSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDockerSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) PutRStudioServerProDomainSettings(value *TfDomain_RStudioServerProDomainSettingsProperty) {
	if err := t.validatePutRStudioServerProDomainSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRStudioServerProDomainSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) PutTrustedIdentityPropagationSettings(value *TfDomain_TrustedIdentityPropagationSettingsProperty) {
	if err := t.validatePutTrustedIdentityPropagationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTrustedIdentityPropagationSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) ResetDockerSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetDockerSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) ResetExecutionRoleIdentityConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetExecutionRoleIdentityConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) ResetRStudioServerProDomainSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetRStudioServerProDomainSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) ResetTrustedIdentityPropagationSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetTrustedIdentityPropagationSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDomain_DomainSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

