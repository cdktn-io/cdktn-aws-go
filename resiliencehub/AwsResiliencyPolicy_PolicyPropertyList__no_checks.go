//go:build no_runtime_type_checking

package resiliencehub

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsResiliencyPolicy_PolicyPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsResiliencyPolicy_PolicyPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

