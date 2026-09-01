package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEcsService_DeploymentConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BakeTimeInMinutes() *string
	// Experimental.
	SetBakeTimeInMinutes(val *string)
	// Experimental.
	BakeTimeInMinutesInput() *string
	// Experimental.
	CanaryConfiguration() AwsEcsService_CanaryConfigurationPropertyOutputReference
	// Experimental.
	CanaryConfigurationInput() *AwsEcsService_CanaryConfigurationProperty
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
	InternalValue() *AwsEcsService_DeploymentConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsEcsService_DeploymentConfigurationProperty)
	// Experimental.
	LifecycleHook() AwsEcsService_LifecycleHookPropertyList
	// Experimental.
	LifecycleHookInput() interface{}
	// Experimental.
	LinearConfiguration() AwsEcsService_LinearConfigurationPropertyOutputReference
	// Experimental.
	LinearConfigurationInput() *AwsEcsService_LinearConfigurationProperty
	// Experimental.
	Strategy() *string
	// Experimental.
	SetStrategy(val *string)
	// Experimental.
	StrategyInput() *string
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
	PutCanaryConfiguration(value *AwsEcsService_CanaryConfigurationProperty)
	// Experimental.
	PutLifecycleHook(value interface{})
	// Experimental.
	PutLinearConfiguration(value *AwsEcsService_LinearConfigurationProperty)
	// Experimental.
	ResetBakeTimeInMinutes()
	// Experimental.
	ResetCanaryConfiguration()
	// Experimental.
	ResetLifecycleHook()
	// Experimental.
	ResetLinearConfiguration()
	// Experimental.
	ResetStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEcsService_DeploymentConfigurationPropertyOutputReference
type jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) BakeTimeInMinutes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bakeTimeInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) BakeTimeInMinutesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bakeTimeInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) CanaryConfiguration() AwsEcsService_CanaryConfigurationPropertyOutputReference {
	var returns AwsEcsService_CanaryConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"canaryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) CanaryConfigurationInput() *AwsEcsService_CanaryConfigurationProperty {
	var returns *AwsEcsService_CanaryConfigurationProperty
	_jsii_.Get(
		j,
		"canaryConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) InternalValue() *AwsEcsService_DeploymentConfigurationProperty {
	var returns *AwsEcsService_DeploymentConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) LifecycleHook() AwsEcsService_LifecycleHookPropertyList {
	var returns AwsEcsService_LifecycleHookPropertyList
	_jsii_.Get(
		j,
		"lifecycleHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) LifecycleHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) LinearConfiguration() AwsEcsService_LinearConfigurationPropertyOutputReference {
	var returns AwsEcsService_LinearConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"linearConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) LinearConfigurationInput() *AwsEcsService_LinearConfigurationProperty {
	var returns *AwsEcsService_LinearConfigurationProperty
	_jsii_.Get(
		j,
		"linearConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) Strategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) StrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEcsService_DeploymentConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEcsService_DeploymentConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEcsService_DeploymentConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsEcsService.DeploymentConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEcsService_DeploymentConfigurationPropertyOutputReference_Override(a AwsEcsService_DeploymentConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsEcsService.DeploymentConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference)SetBakeTimeInMinutes(val *string) {
	if err := j.validateSetBakeTimeInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bakeTimeInMinutes",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference)SetInternalValue(val *AwsEcsService_DeploymentConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference)SetStrategy(val *string) {
	if err := j.validateSetStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategy",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) PutCanaryConfiguration(value *AwsEcsService_CanaryConfigurationProperty) {
	if err := a.validatePutCanaryConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCanaryConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) PutLifecycleHook(value interface{}) {
	if err := a.validatePutLifecycleHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLifecycleHook",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) PutLinearConfiguration(value *AwsEcsService_LinearConfigurationProperty) {
	if err := a.validatePutLinearConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLinearConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) ResetBakeTimeInMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetBakeTimeInMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) ResetCanaryConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCanaryConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) ResetLifecycleHook() {
	_jsii_.InvokeVoid(
		a,
		"resetLifecycleHook",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) ResetLinearConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLinearConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) ResetStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEcsService_DeploymentConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

