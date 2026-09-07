package guardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/guardduty/jsii"

	"github.com/cdktn-io/cdktn-aws-go/guardduty/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsOrganizationConfiguration_DatasourcesPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() *AwsOrganizationConfiguration_DatasourcesProperty
	// Experimental.
	SetInternalValue(val *AwsOrganizationConfiguration_DatasourcesProperty)
	// Experimental.
	Kubernetes() AwsOrganizationConfiguration_KubernetesPropertyOutputReference
	// Experimental.
	KubernetesInput() *AwsOrganizationConfiguration_KubernetesProperty
	// Experimental.
	MalwareProtection() AwsOrganizationConfiguration_MalwareProtectionPropertyOutputReference
	// Experimental.
	MalwareProtectionInput() *AwsOrganizationConfiguration_MalwareProtectionProperty
	// Experimental.
	S3Logs() AwsOrganizationConfiguration_S3LogsPropertyOutputReference
	// Experimental.
	S3LogsInput() *AwsOrganizationConfiguration_S3LogsProperty
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
	PutKubernetes(value *AwsOrganizationConfiguration_KubernetesProperty)
	// Experimental.
	PutMalwareProtection(value *AwsOrganizationConfiguration_MalwareProtectionProperty)
	// Experimental.
	PutS3Logs(value *AwsOrganizationConfiguration_S3LogsProperty)
	// Experimental.
	ResetKubernetes()
	// Experimental.
	ResetMalwareProtection()
	// Experimental.
	ResetS3Logs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsOrganizationConfiguration_DatasourcesPropertyOutputReference
type jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) InternalValue() *AwsOrganizationConfiguration_DatasourcesProperty {
	var returns *AwsOrganizationConfiguration_DatasourcesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) Kubernetes() AwsOrganizationConfiguration_KubernetesPropertyOutputReference {
	var returns AwsOrganizationConfiguration_KubernetesPropertyOutputReference
	_jsii_.Get(
		j,
		"kubernetes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) KubernetesInput() *AwsOrganizationConfiguration_KubernetesProperty {
	var returns *AwsOrganizationConfiguration_KubernetesProperty
	_jsii_.Get(
		j,
		"kubernetesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) MalwareProtection() AwsOrganizationConfiguration_MalwareProtectionPropertyOutputReference {
	var returns AwsOrganizationConfiguration_MalwareProtectionPropertyOutputReference
	_jsii_.Get(
		j,
		"malwareProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) MalwareProtectionInput() *AwsOrganizationConfiguration_MalwareProtectionProperty {
	var returns *AwsOrganizationConfiguration_MalwareProtectionProperty
	_jsii_.Get(
		j,
		"malwareProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) S3Logs() AwsOrganizationConfiguration_S3LogsPropertyOutputReference {
	var returns AwsOrganizationConfiguration_S3LogsPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) S3LogsInput() *AwsOrganizationConfiguration_S3LogsProperty {
	var returns *AwsOrganizationConfiguration_S3LogsProperty
	_jsii_.Get(
		j,
		"s3LogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsOrganizationConfiguration_DatasourcesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsOrganizationConfiguration_DatasourcesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsOrganizationConfiguration_DatasourcesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-guardduty.AwsOrganizationConfiguration.DatasourcesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsOrganizationConfiguration_DatasourcesPropertyOutputReference_Override(a AwsOrganizationConfiguration_DatasourcesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-guardduty.AwsOrganizationConfiguration.DatasourcesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference)SetInternalValue(val *AwsOrganizationConfiguration_DatasourcesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) PutKubernetes(value *AwsOrganizationConfiguration_KubernetesProperty) {
	if err := a.validatePutKubernetesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKubernetes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) PutMalwareProtection(value *AwsOrganizationConfiguration_MalwareProtectionProperty) {
	if err := a.validatePutMalwareProtectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMalwareProtection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) PutS3Logs(value *AwsOrganizationConfiguration_S3LogsProperty) {
	if err := a.validatePutS3LogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Logs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) ResetKubernetes() {
	_jsii_.InvokeVoid(
		a,
		"resetKubernetes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) ResetMalwareProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetMalwareProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) ResetS3Logs() {
	_jsii_.InvokeVoid(
		a,
		"resetS3Logs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsOrganizationConfiguration_DatasourcesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

