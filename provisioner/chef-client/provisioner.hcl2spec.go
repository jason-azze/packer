// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package chefclient

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName            *string                `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType          *string                `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug                *bool                  `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce                *bool                  `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError              *string                `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars             map[string]string      `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars        []string               `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	Json                       map[string]interface{} `cty:"json"`
	ChefEnvironment            *string                `mapstructure:"chef_environment" cty:"chef_environment"`
	ChefLicense                *string                `mapstructure:"chef_license" cty:"chef_license"`
	ClientKey                  *string                `mapstructure:"client_key" cty:"client_key"`
	ConfigTemplate             *string                `mapstructure:"config_template" cty:"config_template"`
	ElevatedUser               *string                `mapstructure:"elevated_user" cty:"elevated_user"`
	ElevatedPassword           *string                `mapstructure:"elevated_password" cty:"elevated_password"`
	EncryptedDataBagSecretPath *string                `mapstructure:"encrypted_data_bag_secret_path" cty:"encrypted_data_bag_secret_path"`
	ExecuteCommand             *string                `mapstructure:"execute_command" cty:"execute_command"`
	GuestOSType                *string                `mapstructure:"guest_os_type" cty:"guest_os_type"`
	InstallCommand             *string                `mapstructure:"install_command" cty:"install_command"`
	KnifeCommand               *string                `mapstructure:"knife_command" cty:"knife_command"`
	NodeName                   *string                `mapstructure:"node_name" cty:"node_name"`
	PolicyGroup                *string                `mapstructure:"policy_group" cty:"policy_group"`
	PolicyName                 *string                `mapstructure:"policy_name" cty:"policy_name"`
	PreventSudo                *bool                  `mapstructure:"prevent_sudo" cty:"prevent_sudo"`
	RunList                    []string               `mapstructure:"run_list" cty:"run_list"`
	ServerUrl                  *string                `mapstructure:"server_url" cty:"server_url"`
	SkipCleanClient            *bool                  `mapstructure:"skip_clean_client" cty:"skip_clean_client"`
	SkipCleanNode              *bool                  `mapstructure:"skip_clean_node" cty:"skip_clean_node"`
	SkipCleanStagingDirectory  *bool                  `mapstructure:"skip_clean_staging_directory" cty:"skip_clean_staging_directory"`
	SkipInstall                *bool                  `mapstructure:"skip_install" cty:"skip_install"`
	SslVerifyMode              *string                `mapstructure:"ssl_verify_mode" cty:"ssl_verify_mode"`
	TrustedCertsDir            *string                `mapstructure:"trusted_certs_dir" cty:"trusted_certs_dir"`
	StagingDir                 *string                `mapstructure:"staging_directory" cty:"staging_directory"`
	ValidationClientName       *string                `mapstructure:"validation_client_name" cty:"validation_client_name"`
	ValidationKeyPath          *string                `mapstructure:"validation_key_path" cty:"validation_key_path"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{} { return new(FlatConfig) }

// HCL2Spec returns the hcldec.Spec of a FlatConfig.
// This spec is used by HCL to read the fields of FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":              &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":            &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                   &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                   &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":                &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":          &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables":     &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"json":                           &hcldec.BlockAttrsSpec{TypeName: "json", ElementType: cty.String, Required: false},
		"chef_environment":               &hcldec.AttrSpec{Name: "chef_environment", Type: cty.String, Required: false},
		"chef_license":                   &hcldec.AttrSpec{Name: "chef_license", Type: cty.String, Required: false},
		"client_key":                     &hcldec.AttrSpec{Name: "client_key", Type: cty.String, Required: false},
		"config_template":                &hcldec.AttrSpec{Name: "config_template", Type: cty.String, Required: false},
		"elevated_user":                  &hcldec.AttrSpec{Name: "elevated_user", Type: cty.String, Required: false},
		"elevated_password":              &hcldec.AttrSpec{Name: "elevated_password", Type: cty.String, Required: false},
		"encrypted_data_bag_secret_path": &hcldec.AttrSpec{Name: "encrypted_data_bag_secret_path", Type: cty.String, Required: false},
		"execute_command":                &hcldec.AttrSpec{Name: "execute_command", Type: cty.String, Required: false},
		"guest_os_type":                  &hcldec.AttrSpec{Name: "guest_os_type", Type: cty.String, Required: false},
		"install_command":                &hcldec.AttrSpec{Name: "install_command", Type: cty.String, Required: false},
		"knife_command":                  &hcldec.AttrSpec{Name: "knife_command", Type: cty.String, Required: false},
		"node_name":                      &hcldec.AttrSpec{Name: "node_name", Type: cty.String, Required: false},
		"policy_group":                   &hcldec.AttrSpec{Name: "policy_group", Type: cty.String, Required: false},
		"policy_name":                    &hcldec.AttrSpec{Name: "policy_name", Type: cty.String, Required: false},
		"prevent_sudo":                   &hcldec.AttrSpec{Name: "prevent_sudo", Type: cty.Bool, Required: false},
		"run_list":                       &hcldec.AttrSpec{Name: "run_list", Type: cty.List(cty.String), Required: false},
		"server_url":                     &hcldec.AttrSpec{Name: "server_url", Type: cty.String, Required: false},
		"skip_clean_client":              &hcldec.AttrSpec{Name: "skip_clean_client", Type: cty.Bool, Required: false},
		"skip_clean_node":                &hcldec.AttrSpec{Name: "skip_clean_node", Type: cty.Bool, Required: false},
		"skip_clean_staging_directory":   &hcldec.AttrSpec{Name: "skip_clean_staging_directory", Type: cty.Bool, Required: false},
		"skip_install":                   &hcldec.AttrSpec{Name: "skip_install", Type: cty.Bool, Required: false},
		"ssl_verify_mode":                &hcldec.AttrSpec{Name: "ssl_verify_mode", Type: cty.String, Required: false},
		"trusted_certs_dir":              &hcldec.AttrSpec{Name: "trusted_certs_dir", Type: cty.String, Required: false},
		"staging_directory":              &hcldec.AttrSpec{Name: "staging_directory", Type: cty.String, Required: false},
		"validation_client_name":         &hcldec.AttrSpec{Name: "validation_client_name", Type: cty.String, Required: false},
		"validation_key_path":            &hcldec.AttrSpec{Name: "validation_key_path", Type: cty.String, Required: false},
	}
	return s
}