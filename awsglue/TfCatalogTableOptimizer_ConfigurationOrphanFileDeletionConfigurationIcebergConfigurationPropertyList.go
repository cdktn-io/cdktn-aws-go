package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList interface {
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
	Get(index *float64) TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList
type jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList {
	_init_.Initialize()

	if err := validateNewTfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-glue.TfCatalogTableOptimizer.ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList_Override(t TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.TfCatalogTableOptimizer.ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		t,
	)
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := t.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		t,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList) Get(index *float64) TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference {
	if err := t.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyOutputReference

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCatalogTableOptimizer_ConfigurationOrphanFileDeletionConfigurationIcebergConfigurationPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

