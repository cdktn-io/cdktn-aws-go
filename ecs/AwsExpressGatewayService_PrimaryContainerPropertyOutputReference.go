package ecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/ecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsExpressGatewayService_PrimaryContainerPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AwsLogsConfiguration() AwsExpressGatewayService_AwsLogsConfigurationPropertyList
	// Experimental.
	AwsLogsConfigurationInput() interface{}
	// Experimental.
	Command() *[]*string
	// Experimental.
	SetCommand(val *[]*string)
	// Experimental.
	CommandInput() *[]*string
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
	ContainerPort() *float64
	// Experimental.
	SetContainerPort(val *float64)
	// Experimental.
	ContainerPortInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Environment() AwsExpressGatewayService_EnvironmentPropertyList
	// Experimental.
	EnvironmentInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	Image() *string
	// Experimental.
	SetImage(val *string)
	// Experimental.
	ImageInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	RepositoryCredentials() AwsExpressGatewayService_RepositoryCredentialsPropertyList
	// Experimental.
	RepositoryCredentialsInput() interface{}
	// Experimental.
	Secret() AwsExpressGatewayService_SecretPropertyList
	// Experimental.
	SecretInput() interface{}
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
	PutAwsLogsConfiguration(value interface{})
	// Experimental.
	PutEnvironment(value interface{})
	// Experimental.
	PutRepositoryCredentials(value interface{})
	// Experimental.
	PutSecret(value interface{})
	// Experimental.
	ResetAwsLogsConfiguration()
	// Experimental.
	ResetCommand()
	// Experimental.
	ResetContainerPort()
	// Experimental.
	ResetEnvironment()
	// Experimental.
	ResetRepositoryCredentials()
	// Experimental.
	ResetSecret()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsExpressGatewayService_PrimaryContainerPropertyOutputReference
type jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) AwsLogsConfiguration() AwsExpressGatewayService_AwsLogsConfigurationPropertyList {
	var returns AwsExpressGatewayService_AwsLogsConfigurationPropertyList
	_jsii_.Get(
		j,
		"awsLogsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) AwsLogsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsLogsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) ContainerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) ContainerPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) Environment() AwsExpressGatewayService_EnvironmentPropertyList {
	var returns AwsExpressGatewayService_EnvironmentPropertyList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) RepositoryCredentials() AwsExpressGatewayService_RepositoryCredentialsPropertyList {
	var returns AwsExpressGatewayService_RepositoryCredentialsPropertyList
	_jsii_.Get(
		j,
		"repositoryCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) RepositoryCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) Secret() AwsExpressGatewayService_SecretPropertyList {
	var returns AwsExpressGatewayService_SecretPropertyList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsExpressGatewayService_PrimaryContainerPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsExpressGatewayService_PrimaryContainerPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsExpressGatewayService_PrimaryContainerPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsExpressGatewayService.PrimaryContainerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsExpressGatewayService_PrimaryContainerPropertyOutputReference_Override(a AwsExpressGatewayService_PrimaryContainerPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsExpressGatewayService.PrimaryContainerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference)SetContainerPort(val *float64) {
	if err := j.validateSetContainerPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerPort",
		val,
	)
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) PutAwsLogsConfiguration(value interface{}) {
	if err := a.validatePutAwsLogsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsLogsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) PutEnvironment(value interface{}) {
	if err := a.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) PutRepositoryCredentials(value interface{}) {
	if err := a.validatePutRepositoryCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRepositoryCredentials",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) PutSecret(value interface{}) {
	if err := a.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecret",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) ResetAwsLogsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsLogsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetCommand",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) ResetContainerPort() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) ResetRepositoryCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetRepositoryCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsExpressGatewayService_PrimaryContainerPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

