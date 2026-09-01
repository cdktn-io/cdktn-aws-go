package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseks/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference interface {
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
	InternalValue() *DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsProperty
	// Experimental.
	SetInternalValue(val *DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsProperty)
	// Experimental.
	MaxPort() DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsMaxPortPropertyList
	// Experimental.
	MinPort() DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsMinPortPropertyList
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

// The jsii proxy struct for DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference
type jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) InternalValue() *DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsProperty {
	var returns *DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) MaxPort() DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsMaxPortPropertyList {
	var returns DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsMaxPortPropertyList
	_jsii_.Get(
		j,
		"maxPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) MinPort() DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsMinPortPropertyList {
	var returns DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsMinPortPropertyList
	_jsii_.Get(
		j,
		"minPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eks.DataAwsEksClusterVersions.ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference_Override(d DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.DataAwsEksClusterVersions.ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference)SetInternalValue(val *DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsEksClusterVersions_ClusterVersionsControlPlaneComponentConfigKubeApiServerConfigServiceNodePortRangeConstraintsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

