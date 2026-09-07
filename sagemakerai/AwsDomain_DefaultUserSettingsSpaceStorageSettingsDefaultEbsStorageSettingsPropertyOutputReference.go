package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference interface {
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
	DefaultEbsVolumeSizeInGb() *float64
	// Experimental.
	SetDefaultEbsVolumeSizeInGb(val *float64)
	// Experimental.
	DefaultEbsVolumeSizeInGbInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsProperty)
	// Experimental.
	MaximumEbsVolumeSizeInGb() *float64
	// Experimental.
	SetMaximumEbsVolumeSizeInGb(val *float64)
	// Experimental.
	MaximumEbsVolumeSizeInGbInput() *float64
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

// The jsii proxy struct for AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference
type jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) DefaultEbsVolumeSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultEbsVolumeSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) DefaultEbsVolumeSizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultEbsVolumeSizeInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) InternalValue() *AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsProperty {
	var returns *AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) MaximumEbsVolumeSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumEbsVolumeSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) MaximumEbsVolumeSizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumEbsVolumeSizeInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDomain.DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference_Override(a AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDomain.DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference)SetDefaultEbsVolumeSizeInGb(val *float64) {
	if err := j.validateSetDefaultEbsVolumeSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultEbsVolumeSizeInGb",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference)SetInternalValue(val *AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference)SetMaximumEbsVolumeSizeInGb(val *float64) {
	if err := j.validateSetMaximumEbsVolumeSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumEbsVolumeSizeInGb",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

