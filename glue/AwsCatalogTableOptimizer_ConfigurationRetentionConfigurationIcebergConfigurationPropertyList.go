package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/glue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/glue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList interface {
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
	Get(index *float64) AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList
type jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList {
	_init_.Initialize()

	if err := validateNewAwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsCatalogTableOptimizer.ConfigurationRetentionConfigurationIcebergConfigurationPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList_Override(a AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsCatalogTableOptimizer.ConfigurationRetentionConfigurationIcebergConfigurationPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList) Get(index *float64) AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCatalogTableOptimizer_ConfigurationRetentionConfigurationIcebergConfigurationPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

