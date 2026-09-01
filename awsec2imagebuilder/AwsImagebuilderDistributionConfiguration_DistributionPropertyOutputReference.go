package awsec2imagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AmiDistributionConfiguration() AwsImagebuilderDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference
	// Experimental.
	AmiDistributionConfigurationInput() *AwsImagebuilderDistributionConfiguration_AmiDistributionConfigurationProperty
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
	ContainerDistributionConfiguration() AwsImagebuilderDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference
	// Experimental.
	ContainerDistributionConfigurationInput() *AwsImagebuilderDistributionConfiguration_ContainerDistributionConfigurationProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	FastLaunchConfiguration() AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyList
	// Experimental.
	FastLaunchConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LaunchTemplateConfiguration() AwsImagebuilderDistributionConfiguration_LaunchTemplateConfigurationPropertyList
	// Experimental.
	LaunchTemplateConfigurationInput() interface{}
	// Experimental.
	LicenseConfigurationArns() *[]*string
	// Experimental.
	SetLicenseConfigurationArns(val *[]*string)
	// Experimental.
	LicenseConfigurationArnsInput() *[]*string
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	S3ExportConfiguration() AwsImagebuilderDistributionConfiguration_S3ExportConfigurationPropertyOutputReference
	// Experimental.
	S3ExportConfigurationInput() *AwsImagebuilderDistributionConfiguration_S3ExportConfigurationProperty
	// Experimental.
	SsmParameterConfiguration() AwsImagebuilderDistributionConfiguration_SsmParameterConfigurationPropertyList
	// Experimental.
	SsmParameterConfigurationInput() interface{}
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
	PutAmiDistributionConfiguration(value *AwsImagebuilderDistributionConfiguration_AmiDistributionConfigurationProperty)
	// Experimental.
	PutContainerDistributionConfiguration(value *AwsImagebuilderDistributionConfiguration_ContainerDistributionConfigurationProperty)
	// Experimental.
	PutFastLaunchConfiguration(value interface{})
	// Experimental.
	PutLaunchTemplateConfiguration(value interface{})
	// Experimental.
	PutS3ExportConfiguration(value *AwsImagebuilderDistributionConfiguration_S3ExportConfigurationProperty)
	// Experimental.
	PutSsmParameterConfiguration(value interface{})
	// Experimental.
	ResetAmiDistributionConfiguration()
	// Experimental.
	ResetContainerDistributionConfiguration()
	// Experimental.
	ResetFastLaunchConfiguration()
	// Experimental.
	ResetLaunchTemplateConfiguration()
	// Experimental.
	ResetLicenseConfigurationArns()
	// Experimental.
	ResetS3ExportConfiguration()
	// Experimental.
	ResetSsmParameterConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference
type jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) AmiDistributionConfiguration() AwsImagebuilderDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference {
	var returns AwsImagebuilderDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"amiDistributionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) AmiDistributionConfigurationInput() *AwsImagebuilderDistributionConfiguration_AmiDistributionConfigurationProperty {
	var returns *AwsImagebuilderDistributionConfiguration_AmiDistributionConfigurationProperty
	_jsii_.Get(
		j,
		"amiDistributionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) ContainerDistributionConfiguration() AwsImagebuilderDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference {
	var returns AwsImagebuilderDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"containerDistributionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) ContainerDistributionConfigurationInput() *AwsImagebuilderDistributionConfiguration_ContainerDistributionConfigurationProperty {
	var returns *AwsImagebuilderDistributionConfiguration_ContainerDistributionConfigurationProperty
	_jsii_.Get(
		j,
		"containerDistributionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) FastLaunchConfiguration() AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyList {
	var returns AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyList
	_jsii_.Get(
		j,
		"fastLaunchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) FastLaunchConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fastLaunchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) LaunchTemplateConfiguration() AwsImagebuilderDistributionConfiguration_LaunchTemplateConfigurationPropertyList {
	var returns AwsImagebuilderDistributionConfiguration_LaunchTemplateConfigurationPropertyList
	_jsii_.Get(
		j,
		"launchTemplateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) LaunchTemplateConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) LicenseConfigurationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licenseConfigurationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) LicenseConfigurationArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licenseConfigurationArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) S3ExportConfiguration() AwsImagebuilderDistributionConfiguration_S3ExportConfigurationPropertyOutputReference {
	var returns AwsImagebuilderDistributionConfiguration_S3ExportConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3ExportConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) S3ExportConfigurationInput() *AwsImagebuilderDistributionConfiguration_S3ExportConfigurationProperty {
	var returns *AwsImagebuilderDistributionConfiguration_S3ExportConfigurationProperty
	_jsii_.Get(
		j,
		"s3ExportConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) SsmParameterConfiguration() AwsImagebuilderDistributionConfiguration_SsmParameterConfigurationPropertyList {
	var returns AwsImagebuilderDistributionConfiguration_SsmParameterConfigurationPropertyList
	_jsii_.Get(
		j,
		"ssmParameterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) SsmParameterConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssmParameterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.AwsImagebuilderDistributionConfiguration.DistributionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference_Override(a AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.AwsImagebuilderDistributionConfiguration.DistributionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference)SetLicenseConfigurationArns(val *[]*string) {
	if err := j.validateSetLicenseConfigurationArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseConfigurationArns",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) PutAmiDistributionConfiguration(value *AwsImagebuilderDistributionConfiguration_AmiDistributionConfigurationProperty) {
	if err := a.validatePutAmiDistributionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmiDistributionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) PutContainerDistributionConfiguration(value *AwsImagebuilderDistributionConfiguration_ContainerDistributionConfigurationProperty) {
	if err := a.validatePutContainerDistributionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContainerDistributionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) PutFastLaunchConfiguration(value interface{}) {
	if err := a.validatePutFastLaunchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFastLaunchConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) PutLaunchTemplateConfiguration(value interface{}) {
	if err := a.validatePutLaunchTemplateConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLaunchTemplateConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) PutS3ExportConfiguration(value *AwsImagebuilderDistributionConfiguration_S3ExportConfigurationProperty) {
	if err := a.validatePutS3ExportConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3ExportConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) PutSsmParameterConfiguration(value interface{}) {
	if err := a.validatePutSsmParameterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSsmParameterConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) ResetAmiDistributionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAmiDistributionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) ResetContainerDistributionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerDistributionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) ResetFastLaunchConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetFastLaunchConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) ResetLaunchTemplateConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchTemplateConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) ResetLicenseConfigurationArns() {
	_jsii_.InvokeVoid(
		a,
		"resetLicenseConfigurationArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) ResetS3ExportConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetS3ExportConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) ResetSsmParameterConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSsmParameterConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_DistributionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

