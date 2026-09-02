package awsec2imagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AmiTags() *map[string]*string
	// Experimental.
	SetAmiTags(val *map[string]*string)
	// Experimental.
	AmiTagsInput() *map[string]*string
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
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDistributionConfiguration_AmiDistributionConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfDistributionConfiguration_AmiDistributionConfigurationProperty)
	// Experimental.
	KmsKeyId() *string
	// Experimental.
	SetKmsKeyId(val *string)
	// Experimental.
	KmsKeyIdInput() *string
	// Experimental.
	LaunchPermission() TfDistributionConfiguration_LaunchPermissionPropertyOutputReference
	// Experimental.
	LaunchPermissionInput() *TfDistributionConfiguration_LaunchPermissionProperty
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	TargetAccountIds() *[]*string
	// Experimental.
	SetTargetAccountIds(val *[]*string)
	// Experimental.
	TargetAccountIdsInput() *[]*string
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
	PutLaunchPermission(value *TfDistributionConfiguration_LaunchPermissionProperty)
	// Experimental.
	ResetAmiTags()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetKmsKeyId()
	// Experimental.
	ResetLaunchPermission()
	// Experimental.
	ResetName()
	// Experimental.
	ResetTargetAccountIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference
type jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) AmiTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"amiTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) AmiTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"amiTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) InternalValue() *TfDistributionConfiguration_AmiDistributionConfigurationProperty {
	var returns *TfDistributionConfiguration_AmiDistributionConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) LaunchPermission() TfDistributionConfiguration_LaunchPermissionPropertyOutputReference {
	var returns TfDistributionConfiguration_LaunchPermissionPropertyOutputReference
	_jsii_.Get(
		j,
		"launchPermission",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) LaunchPermissionInput() *TfDistributionConfiguration_LaunchPermissionProperty {
	var returns *TfDistributionConfiguration_LaunchPermissionProperty
	_jsii_.Get(
		j,
		"launchPermissionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) TargetAccountIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetAccountIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) TargetAccountIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetAccountIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.TfDistributionConfiguration.AmiDistributionConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference_Override(t TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.TfDistributionConfiguration.AmiDistributionConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference)SetAmiTags(val *map[string]*string) {
	if err := j.validateSetAmiTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amiTags",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference)SetInternalValue(val *TfDistributionConfiguration_AmiDistributionConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference)SetTargetAccountIds(val *[]*string) {
	if err := j.validateSetTargetAccountIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetAccountIds",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) PutLaunchPermission(value *TfDistributionConfiguration_LaunchPermissionProperty) {
	if err := t.validatePutLaunchPermissionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLaunchPermission",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) ResetAmiTags() {
	_jsii_.InvokeVoid(
		t,
		"resetAmiTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) ResetLaunchPermission() {
	_jsii_.InvokeVoid(
		t,
		"resetLaunchPermission",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) ResetTargetAccountIds() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetAccountIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDistributionConfiguration_AmiDistributionConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

