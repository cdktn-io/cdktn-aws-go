package awschimesdkmediapipelines

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awschimesdkmediapipelines/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awschimesdkmediapipelines/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference interface {
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
	ContentRedactionOutput() *string
	// Experimental.
	SetContentRedactionOutput(val *string)
	// Experimental.
	ContentRedactionOutputInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DataAccessRoleArn() *string
	// Experimental.
	SetDataAccessRoleArn(val *string)
	// Experimental.
	DataAccessRoleArnInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty)
	// Experimental.
	OutputEncryptionKmsKeyId() *string
	// Experimental.
	SetOutputEncryptionKmsKeyId(val *string)
	// Experimental.
	OutputEncryptionKmsKeyIdInput() *string
	// Experimental.
	OutputLocation() *string
	// Experimental.
	SetOutputLocation(val *string)
	// Experimental.
	OutputLocationInput() *string
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
	ResetContentRedactionOutput()
	// Experimental.
	ResetOutputEncryptionKmsKeyId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference
type jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) ContentRedactionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentRedactionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) ContentRedactionOutputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentRedactionOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) DataAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) DataAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) InternalValue() *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty {
	var returns *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) OutputEncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputEncryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) OutputEncryptionKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputEncryptionKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) OutputLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) OutputLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration.PostCallAnalyticsSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference_Override(a AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration.PostCallAnalyticsSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference)SetContentRedactionOutput(val *string) {
	if err := j.validateSetContentRedactionOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentRedactionOutput",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference)SetDataAccessRoleArn(val *string) {
	if err := j.validateSetDataAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference)SetInternalValue(val *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference)SetOutputEncryptionKmsKeyId(val *string) {
	if err := j.validateSetOutputEncryptionKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputEncryptionKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference)SetOutputLocation(val *string) {
	if err := j.validateSetOutputLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputLocation",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) ResetContentRedactionOutput() {
	_jsii_.InvokeVoid(
		a,
		"resetContentRedactionOutput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) ResetOutputEncryptionKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputEncryptionKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

