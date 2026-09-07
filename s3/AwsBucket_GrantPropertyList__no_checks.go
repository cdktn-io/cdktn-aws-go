//go:build no_runtime_type_checking

package s3

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsBucket_GrantPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsBucket_GrantPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsBucket_GrantPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsBucket_GrantPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsBucket_GrantPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsBucket_GrantPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsBucket_GrantPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsBucket_GrantPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

