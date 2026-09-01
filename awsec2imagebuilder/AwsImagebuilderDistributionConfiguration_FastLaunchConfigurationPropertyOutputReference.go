package awsec2imagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccountId() *string
	// Experimental.
	SetAccountId(val *string)
	// Experimental.
	AccountIdInput() *string
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
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LaunchTemplate() AwsImagebuilderDistributionConfiguration_LaunchTemplatePropertyOutputReference
	// Experimental.
	LaunchTemplateInput() *AwsImagebuilderDistributionConfiguration_LaunchTemplateProperty
	// Experimental.
	MaxParallelLaunches() *float64
	// Experimental.
	SetMaxParallelLaunches(val *float64)
	// Experimental.
	MaxParallelLaunchesInput() *float64
	// Experimental.
	SnapshotConfiguration() AwsImagebuilderDistributionConfiguration_SnapshotConfigurationPropertyOutputReference
	// Experimental.
	SnapshotConfigurationInput() *AwsImagebuilderDistributionConfiguration_SnapshotConfigurationProperty
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
	PutLaunchTemplate(value *AwsImagebuilderDistributionConfiguration_LaunchTemplateProperty)
	// Experimental.
	PutSnapshotConfiguration(value *AwsImagebuilderDistributionConfiguration_SnapshotConfigurationProperty)
	// Experimental.
	ResetLaunchTemplate()
	// Experimental.
	ResetMaxParallelLaunches()
	// Experimental.
	ResetSnapshotConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference
type jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) LaunchTemplate() AwsImagebuilderDistributionConfiguration_LaunchTemplatePropertyOutputReference {
	var returns AwsImagebuilderDistributionConfiguration_LaunchTemplatePropertyOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) LaunchTemplateInput() *AwsImagebuilderDistributionConfiguration_LaunchTemplateProperty {
	var returns *AwsImagebuilderDistributionConfiguration_LaunchTemplateProperty
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) MaxParallelLaunches() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelLaunches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) MaxParallelLaunchesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelLaunchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) SnapshotConfiguration() AwsImagebuilderDistributionConfiguration_SnapshotConfigurationPropertyOutputReference {
	var returns AwsImagebuilderDistributionConfiguration_SnapshotConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"snapshotConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) SnapshotConfigurationInput() *AwsImagebuilderDistributionConfiguration_SnapshotConfigurationProperty {
	var returns *AwsImagebuilderDistributionConfiguration_SnapshotConfigurationProperty
	_jsii_.Get(
		j,
		"snapshotConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.AwsImagebuilderDistributionConfiguration.FastLaunchConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference_Override(a AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.AwsImagebuilderDistributionConfiguration.FastLaunchConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference)SetMaxParallelLaunches(val *float64) {
	if err := j.validateSetMaxParallelLaunchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallelLaunches",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) PutLaunchTemplate(value *AwsImagebuilderDistributionConfiguration_LaunchTemplateProperty) {
	if err := a.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) PutSnapshotConfiguration(value *AwsImagebuilderDistributionConfiguration_SnapshotConfigurationProperty) {
	if err := a.validatePutSnapshotConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnapshotConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) ResetMaxParallelLaunches() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxParallelLaunches",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) ResetSnapshotConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsImagebuilderDistributionConfiguration_FastLaunchConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

