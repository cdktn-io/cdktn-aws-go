package awss3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllowedFeatures() *[]*string
	// Experimental.
	SetAllowedFeatures(val *[]*string)
	// Experimental.
	AllowedFeaturesInput() *[]*string
	// Experimental.
	CloudWatchMetricsEnabled() interface{}
	// Experimental.
	SetCloudWatchMetricsEnabled(val interface{})
	// Experimental.
	CloudWatchMetricsEnabledInput() interface{}
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
	InternalValue() *AwsS3ControlObjectLambdaAccessPoint_ConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsS3ControlObjectLambdaAccessPoint_ConfigurationProperty)
	// Experimental.
	SupportingAccessPoint() *string
	// Experimental.
	SetSupportingAccessPoint(val *string)
	// Experimental.
	SupportingAccessPointInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TransformationConfiguration() AwsS3ControlObjectLambdaAccessPoint_TransformationConfigurationPropertyList
	// Experimental.
	TransformationConfigurationInput() interface{}
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
	PutTransformationConfiguration(value interface{})
	// Experimental.
	ResetAllowedFeatures()
	// Experimental.
	ResetCloudWatchMetricsEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference
type jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) AllowedFeatures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) AllowedFeaturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) CloudWatchMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudWatchMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) CloudWatchMetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudWatchMetricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) InternalValue() *AwsS3ControlObjectLambdaAccessPoint_ConfigurationProperty {
	var returns *AwsS3ControlObjectLambdaAccessPoint_ConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) SupportingAccessPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportingAccessPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) SupportingAccessPointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportingAccessPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) TransformationConfiguration() AwsS3ControlObjectLambdaAccessPoint_TransformationConfigurationPropertyList {
	var returns AwsS3ControlObjectLambdaAccessPoint_TransformationConfigurationPropertyList
	_jsii_.Get(
		j,
		"transformationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) TransformationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformationConfigurationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsS3ControlObjectLambdaAccessPoint.ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference_Override(a AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsS3ControlObjectLambdaAccessPoint.ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference)SetAllowedFeatures(val *[]*string) {
	if err := j.validateSetAllowedFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedFeatures",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference)SetCloudWatchMetricsEnabled(val interface{}) {
	if err := j.validateSetCloudWatchMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudWatchMetricsEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference)SetInternalValue(val *AwsS3ControlObjectLambdaAccessPoint_ConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference)SetSupportingAccessPoint(val *string) {
	if err := j.validateSetSupportingAccessPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportingAccessPoint",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) PutTransformationConfiguration(value interface{}) {
	if err := a.validatePutTransformationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTransformationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) ResetAllowedFeatures() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedFeatures",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) ResetCloudWatchMetricsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudWatchMetricsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsS3ControlObjectLambdaAccessPoint_ConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

