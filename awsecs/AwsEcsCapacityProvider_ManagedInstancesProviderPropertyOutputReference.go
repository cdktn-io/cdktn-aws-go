package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InfrastructureOptimization() AwsEcsCapacityProvider_InfrastructureOptimizationPropertyOutputReference
	// Experimental.
	InfrastructureOptimizationInput() *AwsEcsCapacityProvider_InfrastructureOptimizationProperty
	// Experimental.
	InfrastructureRoleArn() *string
	// Experimental.
	SetInfrastructureRoleArn(val *string)
	// Experimental.
	InfrastructureRoleArnInput() *string
	// Experimental.
	InstanceLaunchTemplate() AwsEcsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference
	// Experimental.
	InstanceLaunchTemplateInput() *AwsEcsCapacityProvider_InstanceLaunchTemplateProperty
	// Experimental.
	InternalValue() *AwsEcsCapacityProvider_ManagedInstancesProviderProperty
	// Experimental.
	SetInternalValue(val *AwsEcsCapacityProvider_ManagedInstancesProviderProperty)
	// Experimental.
	PropagateTags() *string
	// Experimental.
	SetPropagateTags(val *string)
	// Experimental.
	PropagateTagsInput() *string
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
	PutInfrastructureOptimization(value *AwsEcsCapacityProvider_InfrastructureOptimizationProperty)
	// Experimental.
	PutInstanceLaunchTemplate(value *AwsEcsCapacityProvider_InstanceLaunchTemplateProperty)
	// Experimental.
	ResetInfrastructureOptimization()
	// Experimental.
	ResetPropagateTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference
type jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InfrastructureOptimization() AwsEcsCapacityProvider_InfrastructureOptimizationPropertyOutputReference {
	var returns AwsEcsCapacityProvider_InfrastructureOptimizationPropertyOutputReference
	_jsii_.Get(
		j,
		"infrastructureOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InfrastructureOptimizationInput() *AwsEcsCapacityProvider_InfrastructureOptimizationProperty {
	var returns *AwsEcsCapacityProvider_InfrastructureOptimizationProperty
	_jsii_.Get(
		j,
		"infrastructureOptimizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InfrastructureRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InfrastructureRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InstanceLaunchTemplate() AwsEcsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference {
	var returns AwsEcsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference
	_jsii_.Get(
		j,
		"instanceLaunchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InstanceLaunchTemplateInput() *AwsEcsCapacityProvider_InstanceLaunchTemplateProperty {
	var returns *AwsEcsCapacityProvider_InstanceLaunchTemplateProperty
	_jsii_.Get(
		j,
		"instanceLaunchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InternalValue() *AwsEcsCapacityProvider_ManagedInstancesProviderProperty {
	var returns *AwsEcsCapacityProvider_ManagedInstancesProviderProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) PropagateTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsEcsCapacityProvider.ManagedInstancesProviderPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference_Override(a AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsEcsCapacityProvider.ManagedInstancesProviderPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetInfrastructureRoleArn(val *string) {
	if err := j.validateSetInfrastructureRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infrastructureRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetInternalValue(val *AwsEcsCapacityProvider_ManagedInstancesProviderProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetPropagateTags(val *string) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) PutInfrastructureOptimization(value *AwsEcsCapacityProvider_InfrastructureOptimizationProperty) {
	if err := a.validatePutInfrastructureOptimizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInfrastructureOptimization",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) PutInstanceLaunchTemplate(value *AwsEcsCapacityProvider_InstanceLaunchTemplateProperty) {
	if err := a.validatePutInstanceLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstanceLaunchTemplate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ResetInfrastructureOptimization() {
	_jsii_.InvokeVoid(
		a,
		"resetInfrastructureOptimization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		a,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

