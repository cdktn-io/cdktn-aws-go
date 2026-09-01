package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbatch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbatch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	Get(index *float64) AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList
type jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList {
	_init_.Initialize()

	if err := validateNewAwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-batch.AwsBatchJobDefinition.EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList_Override(a AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-batch.AwsBatchJobDefinition.EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList) Get(index *float64) AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBatchJobDefinition_EksPropertiesPodPropertiesInitContainersVolumeMountsPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

