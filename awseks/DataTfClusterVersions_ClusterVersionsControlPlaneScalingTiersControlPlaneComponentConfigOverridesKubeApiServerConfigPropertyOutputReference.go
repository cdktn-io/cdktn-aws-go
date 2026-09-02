package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseks/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference interface {
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
	EventTtl() DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigEventTtlPropertyList
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigProperty
	// Experimental.
	SetInternalValue(val *DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigProperty)
	// Experimental.
	ServiceNodePortRange() DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigServiceNodePortRangePropertyList
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

// The jsii proxy struct for DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference
type jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) EventTtl() DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigEventTtlPropertyList {
	var returns DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigEventTtlPropertyList
	_jsii_.Get(
		j,
		"eventTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) InternalValue() *DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigProperty {
	var returns *DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) ServiceNodePortRange() DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigServiceNodePortRangePropertyList {
	var returns DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigServiceNodePortRangePropertyList
	_jsii_.Get(
		j,
		"serviceNodePortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eks.DataTfClusterVersions.ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference_Override(d DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.DataTfClusterVersions.ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference)SetInternalValue(val *DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeApiServerConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

