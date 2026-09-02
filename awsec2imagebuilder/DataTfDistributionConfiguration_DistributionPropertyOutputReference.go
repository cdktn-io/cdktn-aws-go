package awsec2imagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfDistributionConfiguration_DistributionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AmiDistributionConfiguration() DataTfDistributionConfiguration_AmiDistributionConfigurationPropertyList
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
	// Experimental.
	ContainerDistributionConfiguration() DataTfDistributionConfiguration_ContainerDistributionConfigurationPropertyList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	FastLaunchConfiguration() DataTfDistributionConfiguration_FastLaunchConfigurationPropertyList
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *DataTfDistributionConfiguration_DistributionProperty
	// Experimental.
	SetInternalValue(val *DataTfDistributionConfiguration_DistributionProperty)
	// Experimental.
	LaunchTemplateConfiguration() DataTfDistributionConfiguration_LaunchTemplateConfigurationPropertyList
	// Experimental.
	LicenseConfigurationArns() *[]*string
	// Experimental.
	Region() *string
	// Experimental.
	S3ExportConfiguration() DataTfDistributionConfiguration_S3ExportConfigurationPropertyList
	// Experimental.
	SsmParameterConfiguration() DataTfDistributionConfiguration_SsmParameterConfigurationPropertyList
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

// The jsii proxy struct for DataTfDistributionConfiguration_DistributionPropertyOutputReference
type jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) AmiDistributionConfiguration() DataTfDistributionConfiguration_AmiDistributionConfigurationPropertyList {
	var returns DataTfDistributionConfiguration_AmiDistributionConfigurationPropertyList
	_jsii_.Get(
		j,
		"amiDistributionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) ContainerDistributionConfiguration() DataTfDistributionConfiguration_ContainerDistributionConfigurationPropertyList {
	var returns DataTfDistributionConfiguration_ContainerDistributionConfigurationPropertyList
	_jsii_.Get(
		j,
		"containerDistributionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) FastLaunchConfiguration() DataTfDistributionConfiguration_FastLaunchConfigurationPropertyList {
	var returns DataTfDistributionConfiguration_FastLaunchConfigurationPropertyList
	_jsii_.Get(
		j,
		"fastLaunchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) InternalValue() *DataTfDistributionConfiguration_DistributionProperty {
	var returns *DataTfDistributionConfiguration_DistributionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) LaunchTemplateConfiguration() DataTfDistributionConfiguration_LaunchTemplateConfigurationPropertyList {
	var returns DataTfDistributionConfiguration_LaunchTemplateConfigurationPropertyList
	_jsii_.Get(
		j,
		"launchTemplateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) LicenseConfigurationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licenseConfigurationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) S3ExportConfiguration() DataTfDistributionConfiguration_S3ExportConfigurationPropertyList {
	var returns DataTfDistributionConfiguration_S3ExportConfigurationPropertyList
	_jsii_.Get(
		j,
		"s3ExportConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) SsmParameterConfiguration() DataTfDistributionConfiguration_SsmParameterConfigurationPropertyList {
	var returns DataTfDistributionConfiguration_SsmParameterConfigurationPropertyList
	_jsii_.Get(
		j,
		"ssmParameterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfDistributionConfiguration_DistributionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataTfDistributionConfiguration_DistributionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfDistributionConfiguration_DistributionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.DataTfDistributionConfiguration.DistributionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfDistributionConfiguration_DistributionPropertyOutputReference_Override(d DataTfDistributionConfiguration_DistributionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.DataTfDistributionConfiguration.DistributionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference)SetInternalValue(val *DataTfDistributionConfiguration_DistributionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataTfDistributionConfiguration_DistributionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

