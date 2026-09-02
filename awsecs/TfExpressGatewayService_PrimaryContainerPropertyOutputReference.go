package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfExpressGatewayService_PrimaryContainerPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AwsLogsConfiguration() TfExpressGatewayService_AwsLogsConfigurationPropertyList
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
	Environment() TfExpressGatewayService_EnvironmentPropertyList
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
	RepositoryCredentials() TfExpressGatewayService_RepositoryCredentialsPropertyList
	// Experimental.
	RepositoryCredentialsInput() interface{}
	// Experimental.
	Secret() TfExpressGatewayService_SecretPropertyList
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

// The jsii proxy struct for TfExpressGatewayService_PrimaryContainerPropertyOutputReference
type jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) AwsLogsConfiguration() TfExpressGatewayService_AwsLogsConfigurationPropertyList {
	var returns TfExpressGatewayService_AwsLogsConfigurationPropertyList
	_jsii_.Get(
		j,
		"awsLogsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) AwsLogsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsLogsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) ContainerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) ContainerPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) Environment() TfExpressGatewayService_EnvironmentPropertyList {
	var returns TfExpressGatewayService_EnvironmentPropertyList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) RepositoryCredentials() TfExpressGatewayService_RepositoryCredentialsPropertyList {
	var returns TfExpressGatewayService_RepositoryCredentialsPropertyList
	_jsii_.Get(
		j,
		"repositoryCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) RepositoryCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) Secret() TfExpressGatewayService_SecretPropertyList {
	var returns TfExpressGatewayService_SecretPropertyList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfExpressGatewayService_PrimaryContainerPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfExpressGatewayService_PrimaryContainerPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfExpressGatewayService_PrimaryContainerPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.TfExpressGatewayService.PrimaryContainerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfExpressGatewayService_PrimaryContainerPropertyOutputReference_Override(t TfExpressGatewayService_PrimaryContainerPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.TfExpressGatewayService.PrimaryContainerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference)SetContainerPort(val *float64) {
	if err := j.validateSetContainerPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerPort",
		val,
	)
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) PutAwsLogsConfiguration(value interface{}) {
	if err := t.validatePutAwsLogsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAwsLogsConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) PutEnvironment(value interface{}) {
	if err := t.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) PutRepositoryCredentials(value interface{}) {
	if err := t.validatePutRepositoryCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRepositoryCredentials",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) PutSecret(value interface{}) {
	if err := t.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSecret",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) ResetAwsLogsConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetAwsLogsConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		t,
		"resetCommand",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) ResetContainerPort() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		t,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) ResetRepositoryCredentials() {
	_jsii_.InvokeVoid(
		t,
		"resetRepositoryCredentials",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		t,
		"resetSecret",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfExpressGatewayService_PrimaryContainerPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

