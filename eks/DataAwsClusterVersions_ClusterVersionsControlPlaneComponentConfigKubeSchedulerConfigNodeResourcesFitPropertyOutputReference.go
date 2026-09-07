package eks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/eks/jsii"

	"github.com/cdktn-io/cdktn-aws-go/eks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference interface {
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
	InternalValue() *DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitProperty
	// Experimental.
	SetInternalValue(val *DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitProperty)
	// Experimental.
	ScoringStrategy() DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyPropertyList
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

// The jsii proxy struct for DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference
type jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) InternalValue() *DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitProperty {
	var returns *DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) ScoringStrategy() DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyPropertyList {
	var returns DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitScoringStrategyPropertyList
	_jsii_.Get(
		j,
		"scoringStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eks.DataAwsClusterVersions.ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference_Override(d DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.DataAwsClusterVersions.ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference)SetInternalValue(val *DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeSchedulerConfigNodeResourcesFitPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

