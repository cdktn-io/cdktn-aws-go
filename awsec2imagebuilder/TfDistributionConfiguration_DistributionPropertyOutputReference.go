package awsec2imagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDistributionConfiguration_DistributionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AmiDistributionConfiguration() TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference
	// Experimental.
	AmiDistributionConfigurationInput() *TfDistributionConfiguration_AmiDistributionConfigurationProperty
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
	ContainerDistributionConfiguration() TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference
	// Experimental.
	ContainerDistributionConfigurationInput() *TfDistributionConfiguration_ContainerDistributionConfigurationProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	FastLaunchConfiguration() TfDistributionConfiguration_FastLaunchConfigurationPropertyList
	// Experimental.
	FastLaunchConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LaunchTemplateConfiguration() TfDistributionConfiguration_LaunchTemplateConfigurationPropertyList
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
	S3ExportConfiguration() TfDistributionConfiguration_S3ExportConfigurationPropertyOutputReference
	// Experimental.
	S3ExportConfigurationInput() *TfDistributionConfiguration_S3ExportConfigurationProperty
	// Experimental.
	SsmParameterConfiguration() TfDistributionConfiguration_SsmParameterConfigurationPropertyList
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
	PutAmiDistributionConfiguration(value *TfDistributionConfiguration_AmiDistributionConfigurationProperty)
	// Experimental.
	PutContainerDistributionConfiguration(value *TfDistributionConfiguration_ContainerDistributionConfigurationProperty)
	// Experimental.
	PutFastLaunchConfiguration(value interface{})
	// Experimental.
	PutLaunchTemplateConfiguration(value interface{})
	// Experimental.
	PutS3ExportConfiguration(value *TfDistributionConfiguration_S3ExportConfigurationProperty)
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

// The jsii proxy struct for TfDistributionConfiguration_DistributionPropertyOutputReference
type jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) AmiDistributionConfiguration() TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference {
	var returns TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"amiDistributionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) AmiDistributionConfigurationInput() *TfDistributionConfiguration_AmiDistributionConfigurationProperty {
	var returns *TfDistributionConfiguration_AmiDistributionConfigurationProperty
	_jsii_.Get(
		j,
		"amiDistributionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) ContainerDistributionConfiguration() TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference {
	var returns TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"containerDistributionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) ContainerDistributionConfigurationInput() *TfDistributionConfiguration_ContainerDistributionConfigurationProperty {
	var returns *TfDistributionConfiguration_ContainerDistributionConfigurationProperty
	_jsii_.Get(
		j,
		"containerDistributionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) FastLaunchConfiguration() TfDistributionConfiguration_FastLaunchConfigurationPropertyList {
	var returns TfDistributionConfiguration_FastLaunchConfigurationPropertyList
	_jsii_.Get(
		j,
		"fastLaunchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) FastLaunchConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fastLaunchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) LaunchTemplateConfiguration() TfDistributionConfiguration_LaunchTemplateConfigurationPropertyList {
	var returns TfDistributionConfiguration_LaunchTemplateConfigurationPropertyList
	_jsii_.Get(
		j,
		"launchTemplateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) LaunchTemplateConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) LicenseConfigurationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licenseConfigurationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) LicenseConfigurationArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licenseConfigurationArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) S3ExportConfiguration() TfDistributionConfiguration_S3ExportConfigurationPropertyOutputReference {
	var returns TfDistributionConfiguration_S3ExportConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3ExportConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) S3ExportConfigurationInput() *TfDistributionConfiguration_S3ExportConfigurationProperty {
	var returns *TfDistributionConfiguration_S3ExportConfigurationProperty
	_jsii_.Get(
		j,
		"s3ExportConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) SsmParameterConfiguration() TfDistributionConfiguration_SsmParameterConfigurationPropertyList {
	var returns TfDistributionConfiguration_SsmParameterConfigurationPropertyList
	_jsii_.Get(
		j,
		"ssmParameterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) SsmParameterConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssmParameterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDistributionConfiguration_DistributionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfDistributionConfiguration_DistributionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDistributionConfiguration_DistributionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.TfDistributionConfiguration.DistributionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDistributionConfiguration_DistributionPropertyOutputReference_Override(t TfDistributionConfiguration_DistributionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.TfDistributionConfiguration.DistributionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference)SetLicenseConfigurationArns(val *[]*string) {
	if err := j.validateSetLicenseConfigurationArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseConfigurationArns",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) PutAmiDistributionConfiguration(value *TfDistributionConfiguration_AmiDistributionConfigurationProperty) {
	if err := t.validatePutAmiDistributionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAmiDistributionConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) PutContainerDistributionConfiguration(value *TfDistributionConfiguration_ContainerDistributionConfigurationProperty) {
	if err := t.validatePutContainerDistributionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putContainerDistributionConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) PutFastLaunchConfiguration(value interface{}) {
	if err := t.validatePutFastLaunchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFastLaunchConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) PutLaunchTemplateConfiguration(value interface{}) {
	if err := t.validatePutLaunchTemplateConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLaunchTemplateConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) PutS3ExportConfiguration(value *TfDistributionConfiguration_S3ExportConfigurationProperty) {
	if err := t.validatePutS3ExportConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3ExportConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) PutSsmParameterConfiguration(value interface{}) {
	if err := t.validatePutSsmParameterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSsmParameterConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) ResetAmiDistributionConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetAmiDistributionConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) ResetContainerDistributionConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerDistributionConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) ResetFastLaunchConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetFastLaunchConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) ResetLaunchTemplateConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetLaunchTemplateConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) ResetLicenseConfigurationArns() {
	_jsii_.InvokeVoid(
		t,
		"resetLicenseConfigurationArns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) ResetS3ExportConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetS3ExportConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) ResetSsmParameterConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSsmParameterConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDistributionConfiguration_DistributionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

