package awsec2imagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference interface {
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
	// Experimental.
	ContainerTags() *[]*string
	// Experimental.
	SetContainerTags(val *[]*string)
	// Experimental.
	ContainerTagsInput() *[]*string
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
	InternalValue() *TfDistributionConfiguration_ContainerDistributionConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfDistributionConfiguration_ContainerDistributionConfigurationProperty)
	// Experimental.
	TargetRepository() TfDistributionConfiguration_TargetRepositoryPropertyOutputReference
	// Experimental.
	TargetRepositoryInput() *TfDistributionConfiguration_TargetRepositoryProperty
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
	PutTargetRepository(value *TfDistributionConfiguration_TargetRepositoryProperty)
	// Experimental.
	ResetContainerTags()
	// Experimental.
	ResetDescription()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference
type jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) ContainerTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) ContainerTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) InternalValue() *TfDistributionConfiguration_ContainerDistributionConfigurationProperty {
	var returns *TfDistributionConfiguration_ContainerDistributionConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) TargetRepository() TfDistributionConfiguration_TargetRepositoryPropertyOutputReference {
	var returns TfDistributionConfiguration_TargetRepositoryPropertyOutputReference
	_jsii_.Get(
		j,
		"targetRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) TargetRepositoryInput() *TfDistributionConfiguration_TargetRepositoryProperty {
	var returns *TfDistributionConfiguration_TargetRepositoryProperty
	_jsii_.Get(
		j,
		"targetRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.TfDistributionConfiguration.ContainerDistributionConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference_Override(t TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.TfDistributionConfiguration.ContainerDistributionConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference)SetContainerTags(val *[]*string) {
	if err := j.validateSetContainerTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerTags",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference)SetInternalValue(val *TfDistributionConfiguration_ContainerDistributionConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) PutTargetRepository(value *TfDistributionConfiguration_TargetRepositoryProperty) {
	if err := t.validatePutTargetRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTargetRepository",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) ResetContainerTags() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDistributionConfiguration_ContainerDistributionConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

