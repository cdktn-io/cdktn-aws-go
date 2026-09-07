package codedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/codedeploy/jsii"

	"github.com/cdktn-io/cdktn-aws-go/codedeploy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference interface {
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
	DeploymentReadyOption() AwsDeploymentGroup_DeploymentReadyOptionPropertyOutputReference
	// Experimental.
	DeploymentReadyOptionInput() *AwsDeploymentGroup_DeploymentReadyOptionProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	GreenFleetProvisioningOption() AwsDeploymentGroup_GreenFleetProvisioningOptionPropertyOutputReference
	// Experimental.
	GreenFleetProvisioningOptionInput() *AwsDeploymentGroup_GreenFleetProvisioningOptionProperty
	// Experimental.
	InternalValue() *AwsDeploymentGroup_BlueGreenDeploymentConfigProperty
	// Experimental.
	SetInternalValue(val *AwsDeploymentGroup_BlueGreenDeploymentConfigProperty)
	// Experimental.
	TerminateBlueInstancesOnDeploymentSuccess() AwsDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessPropertyOutputReference
	// Experimental.
	TerminateBlueInstancesOnDeploymentSuccessInput() *AwsDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessProperty
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
	PutDeploymentReadyOption(value *AwsDeploymentGroup_DeploymentReadyOptionProperty)
	// Experimental.
	PutGreenFleetProvisioningOption(value *AwsDeploymentGroup_GreenFleetProvisioningOptionProperty)
	// Experimental.
	PutTerminateBlueInstancesOnDeploymentSuccess(value *AwsDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessProperty)
	// Experimental.
	ResetDeploymentReadyOption()
	// Experimental.
	ResetGreenFleetProvisioningOption()
	// Experimental.
	ResetTerminateBlueInstancesOnDeploymentSuccess()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference
type jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) DeploymentReadyOption() AwsDeploymentGroup_DeploymentReadyOptionPropertyOutputReference {
	var returns AwsDeploymentGroup_DeploymentReadyOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"deploymentReadyOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) DeploymentReadyOptionInput() *AwsDeploymentGroup_DeploymentReadyOptionProperty {
	var returns *AwsDeploymentGroup_DeploymentReadyOptionProperty
	_jsii_.Get(
		j,
		"deploymentReadyOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GreenFleetProvisioningOption() AwsDeploymentGroup_GreenFleetProvisioningOptionPropertyOutputReference {
	var returns AwsDeploymentGroup_GreenFleetProvisioningOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"greenFleetProvisioningOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GreenFleetProvisioningOptionInput() *AwsDeploymentGroup_GreenFleetProvisioningOptionProperty {
	var returns *AwsDeploymentGroup_GreenFleetProvisioningOptionProperty
	_jsii_.Get(
		j,
		"greenFleetProvisioningOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) InternalValue() *AwsDeploymentGroup_BlueGreenDeploymentConfigProperty {
	var returns *AwsDeploymentGroup_BlueGreenDeploymentConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) TerminateBlueInstancesOnDeploymentSuccess() AwsDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessPropertyOutputReference {
	var returns AwsDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessPropertyOutputReference
	_jsii_.Get(
		j,
		"terminateBlueInstancesOnDeploymentSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) TerminateBlueInstancesOnDeploymentSuccessInput() *AwsDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessProperty {
	var returns *AwsDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessProperty
	_jsii_.Get(
		j,
		"terminateBlueInstancesOnDeploymentSuccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codedeploy.AwsDeploymentGroup.BlueGreenDeploymentConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference_Override(a AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codedeploy.AwsDeploymentGroup.BlueGreenDeploymentConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference)SetInternalValue(val *AwsDeploymentGroup_BlueGreenDeploymentConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) PutDeploymentReadyOption(value *AwsDeploymentGroup_DeploymentReadyOptionProperty) {
	if err := a.validatePutDeploymentReadyOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeploymentReadyOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) PutGreenFleetProvisioningOption(value *AwsDeploymentGroup_GreenFleetProvisioningOptionProperty) {
	if err := a.validatePutGreenFleetProvisioningOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGreenFleetProvisioningOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) PutTerminateBlueInstancesOnDeploymentSuccess(value *AwsDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessProperty) {
	if err := a.validatePutTerminateBlueInstancesOnDeploymentSuccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTerminateBlueInstancesOnDeploymentSuccess",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) ResetDeploymentReadyOption() {
	_jsii_.InvokeVoid(
		a,
		"resetDeploymentReadyOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) ResetGreenFleetProvisioningOption() {
	_jsii_.InvokeVoid(
		a,
		"resetGreenFleetProvisioningOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) ResetTerminateBlueInstancesOnDeploymentSuccess() {
	_jsii_.InvokeVoid(
		a,
		"resetTerminateBlueInstancesOnDeploymentSuccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

