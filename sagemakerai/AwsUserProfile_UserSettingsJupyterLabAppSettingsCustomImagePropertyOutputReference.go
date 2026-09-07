package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AppImageConfigName() *string
	// Experimental.
	SetAppImageConfigName(val *string)
	// Experimental.
	AppImageConfigNameInput() *string
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
	ImageName() *string
	// Experimental.
	SetImageName(val *string)
	// Experimental.
	ImageNameInput() *string
	// Experimental.
	ImageVersionNumber() *float64
	// Experimental.
	SetImageVersionNumber(val *float64)
	// Experimental.
	ImageVersionNumberInput() *float64
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
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
	ResetImageVersionNumber()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference
type jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) AppImageConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appImageConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) AppImageConfigNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appImageConfigNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) ImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) ImageVersionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageVersionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) ImageVersionNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageVersionNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsUserProfile.UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference_Override(a AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsUserProfile.UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference)SetAppImageConfigName(val *string) {
	if err := j.validateSetAppImageConfigNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appImageConfigName",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference)SetImageName(val *string) {
	if err := j.validateSetImageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference)SetImageVersionNumber(val *float64) {
	if err := j.validateSetImageVersionNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageVersionNumber",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) ResetImageVersionNumber() {
	_jsii_.InvokeVoid(
		a,
		"resetImageVersionNumber",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

