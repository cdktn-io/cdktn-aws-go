package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerDomain_DomainSettingsPropertyOutputReference interface {
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
	DockerSettings() AwsSagemakerDomain_DockerSettingsPropertyOutputReference
	// Experimental.
	DockerSettingsInput() *AwsSagemakerDomain_DockerSettingsProperty
	// Experimental.
	ExecutionRoleIdentityConfig() *string
	// Experimental.
	SetExecutionRoleIdentityConfig(val *string)
	// Experimental.
	ExecutionRoleIdentityConfigInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsSagemakerDomain_DomainSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerDomain_DomainSettingsProperty)
	// Experimental.
	RStudioServerProDomainSettings() AwsSagemakerDomain_RStudioServerProDomainSettingsPropertyOutputReference
	// Experimental.
	RStudioServerProDomainSettingsInput() *AwsSagemakerDomain_RStudioServerProDomainSettingsProperty
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
	TrustedIdentityPropagationSettings() AwsSagemakerDomain_TrustedIdentityPropagationSettingsPropertyOutputReference
	// Experimental.
	TrustedIdentityPropagationSettingsInput() *AwsSagemakerDomain_TrustedIdentityPropagationSettingsProperty
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
	PutDockerSettings(value *AwsSagemakerDomain_DockerSettingsProperty)
	// Experimental.
	PutRStudioServerProDomainSettings(value *AwsSagemakerDomain_RStudioServerProDomainSettingsProperty)
	// Experimental.
	PutTrustedIdentityPropagationSettings(value *AwsSagemakerDomain_TrustedIdentityPropagationSettingsProperty)
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

// The jsii proxy struct for AwsSagemakerDomain_DomainSettingsPropertyOutputReference
type jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) DockerSettings() AwsSagemakerDomain_DockerSettingsPropertyOutputReference {
	var returns AwsSagemakerDomain_DockerSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"dockerSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) DockerSettingsInput() *AwsSagemakerDomain_DockerSettingsProperty {
	var returns *AwsSagemakerDomain_DockerSettingsProperty
	_jsii_.Get(
		j,
		"dockerSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) ExecutionRoleIdentityConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleIdentityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) ExecutionRoleIdentityConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleIdentityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) InternalValue() *AwsSagemakerDomain_DomainSettingsProperty {
	var returns *AwsSagemakerDomain_DomainSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) RStudioServerProDomainSettings() AwsSagemakerDomain_RStudioServerProDomainSettingsPropertyOutputReference {
	var returns AwsSagemakerDomain_RStudioServerProDomainSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rStudioServerProDomainSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) RStudioServerProDomainSettingsInput() *AwsSagemakerDomain_RStudioServerProDomainSettingsProperty {
	var returns *AwsSagemakerDomain_RStudioServerProDomainSettingsProperty
	_jsii_.Get(
		j,
		"rStudioServerProDomainSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) TrustedIdentityPropagationSettings() AwsSagemakerDomain_TrustedIdentityPropagationSettingsPropertyOutputReference {
	var returns AwsSagemakerDomain_TrustedIdentityPropagationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"trustedIdentityPropagationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) TrustedIdentityPropagationSettingsInput() *AwsSagemakerDomain_TrustedIdentityPropagationSettingsProperty {
	var returns *AwsSagemakerDomain_TrustedIdentityPropagationSettingsProperty
	_jsii_.Get(
		j,
		"trustedIdentityPropagationSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerDomain_DomainSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerDomain_DomainSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerDomain_DomainSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerDomain.DomainSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerDomain_DomainSettingsPropertyOutputReference_Override(a AwsSagemakerDomain_DomainSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerDomain.DomainSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference)SetExecutionRoleIdentityConfig(val *string) {
	if err := j.validateSetExecutionRoleIdentityConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleIdentityConfig",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference)SetInternalValue(val *AwsSagemakerDomain_DomainSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) PutDockerSettings(value *AwsSagemakerDomain_DockerSettingsProperty) {
	if err := a.validatePutDockerSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDockerSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) PutRStudioServerProDomainSettings(value *AwsSagemakerDomain_RStudioServerProDomainSettingsProperty) {
	if err := a.validatePutRStudioServerProDomainSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRStudioServerProDomainSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) PutTrustedIdentityPropagationSettings(value *AwsSagemakerDomain_TrustedIdentityPropagationSettingsProperty) {
	if err := a.validatePutTrustedIdentityPropagationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrustedIdentityPropagationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) ResetDockerSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetDockerSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) ResetExecutionRoleIdentityConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionRoleIdentityConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) ResetRStudioServerProDomainSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRStudioServerProDomainSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) ResetTrustedIdentityPropagationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetTrustedIdentityPropagationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerDomain_DomainSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

