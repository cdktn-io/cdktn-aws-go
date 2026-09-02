package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseks/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference interface {
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
	InternalValue() *DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsProperty
	// Experimental.
	SetInternalValue(val *DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsProperty)
	// Experimental.
	Resources() DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsResourcesPropertyList
	// Experimental.
	ScoringStrategy() DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsScoringStrategyPropertyList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference
type jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) InternalValue() *DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsProperty {
	var returns *DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) Resources() DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsResourcesPropertyList {
	var returns DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsResourcesPropertyList
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) ScoringStrategy() DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsScoringStrategyPropertyList {
	var returns DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsScoringStrategyPropertyList
	_jsii_.Get(
		j,
		"scoringStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eks.DataTfClusterVersions.ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference_Override(d DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.DataTfClusterVersions.ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference)SetInternalValue(val *DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyConstraintsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

