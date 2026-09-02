package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodedeploy/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodedeploy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference interface {
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
	DeploymentReadyOption() TfDeploymentGroup_DeploymentReadyOptionPropertyOutputReference
	// Experimental.
	DeploymentReadyOptionInput() *TfDeploymentGroup_DeploymentReadyOptionProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	GreenFleetProvisioningOption() TfDeploymentGroup_GreenFleetProvisioningOptionPropertyOutputReference
	// Experimental.
	GreenFleetProvisioningOptionInput() *TfDeploymentGroup_GreenFleetProvisioningOptionProperty
	// Experimental.
	InternalValue() *TfDeploymentGroup_BlueGreenDeploymentConfigProperty
	// Experimental.
	SetInternalValue(val *TfDeploymentGroup_BlueGreenDeploymentConfigProperty)
	// Experimental.
	TerminateBlueInstancesOnDeploymentSuccess() TfDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessPropertyOutputReference
	// Experimental.
	TerminateBlueInstancesOnDeploymentSuccessInput() *TfDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessProperty
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
	PutDeploymentReadyOption(value *TfDeploymentGroup_DeploymentReadyOptionProperty)
	// Experimental.
	PutGreenFleetProvisioningOption(value *TfDeploymentGroup_GreenFleetProvisioningOptionProperty)
	// Experimental.
	PutTerminateBlueInstancesOnDeploymentSuccess(value *TfDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessProperty)
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

// The jsii proxy struct for TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference
type jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) DeploymentReadyOption() TfDeploymentGroup_DeploymentReadyOptionPropertyOutputReference {
	var returns TfDeploymentGroup_DeploymentReadyOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"deploymentReadyOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) DeploymentReadyOptionInput() *TfDeploymentGroup_DeploymentReadyOptionProperty {
	var returns *TfDeploymentGroup_DeploymentReadyOptionProperty
	_jsii_.Get(
		j,
		"deploymentReadyOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GreenFleetProvisioningOption() TfDeploymentGroup_GreenFleetProvisioningOptionPropertyOutputReference {
	var returns TfDeploymentGroup_GreenFleetProvisioningOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"greenFleetProvisioningOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GreenFleetProvisioningOptionInput() *TfDeploymentGroup_GreenFleetProvisioningOptionProperty {
	var returns *TfDeploymentGroup_GreenFleetProvisioningOptionProperty
	_jsii_.Get(
		j,
		"greenFleetProvisioningOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) InternalValue() *TfDeploymentGroup_BlueGreenDeploymentConfigProperty {
	var returns *TfDeploymentGroup_BlueGreenDeploymentConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) TerminateBlueInstancesOnDeploymentSuccess() TfDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessPropertyOutputReference {
	var returns TfDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessPropertyOutputReference
	_jsii_.Get(
		j,
		"terminateBlueInstancesOnDeploymentSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) TerminateBlueInstancesOnDeploymentSuccessInput() *TfDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessProperty {
	var returns *TfDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessProperty
	_jsii_.Get(
		j,
		"terminateBlueInstancesOnDeploymentSuccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codedeploy.TfDeploymentGroup.BlueGreenDeploymentConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference_Override(t TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codedeploy.TfDeploymentGroup.BlueGreenDeploymentConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference)SetInternalValue(val *TfDeploymentGroup_BlueGreenDeploymentConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) PutDeploymentReadyOption(value *TfDeploymentGroup_DeploymentReadyOptionProperty) {
	if err := t.validatePutDeploymentReadyOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeploymentReadyOption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) PutGreenFleetProvisioningOption(value *TfDeploymentGroup_GreenFleetProvisioningOptionProperty) {
	if err := t.validatePutGreenFleetProvisioningOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGreenFleetProvisioningOption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) PutTerminateBlueInstancesOnDeploymentSuccess(value *TfDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessProperty) {
	if err := t.validatePutTerminateBlueInstancesOnDeploymentSuccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTerminateBlueInstancesOnDeploymentSuccess",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) ResetDeploymentReadyOption() {
	_jsii_.InvokeVoid(
		t,
		"resetDeploymentReadyOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) ResetGreenFleetProvisioningOption() {
	_jsii_.InvokeVoid(
		t,
		"resetGreenFleetProvisioningOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) ResetTerminateBlueInstancesOnDeploymentSuccess() {
	_jsii_.InvokeVoid(
		t,
		"resetTerminateBlueInstancesOnDeploymentSuccess",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

