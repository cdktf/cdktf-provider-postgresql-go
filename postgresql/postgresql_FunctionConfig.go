// Prebuilt postgresql Provider for Terraform CDK (cdktf)
package postgresql

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Body of the function.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#body Function#body}
	Body *string `field:"required" json:"body" yaml:"body"`
	// Name of the function.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#name Function#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// arg block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#arg Function#arg}
	Arg interface{} `field:"optional" json:"arg" yaml:"arg"`
	// The database where the function is located. If not specified, the provider default database is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#database Function#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Automatically drop objects that depend on the function (such as operators or triggers), and in turn all objects that depend on those objects.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#drop_cascade Function#drop_cascade}
	DropCascade interface{} `field:"optional" json:"dropCascade" yaml:"dropCascade"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#id Function#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Function return type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#returns Function#returns}
	Returns *string `field:"optional" json:"returns" yaml:"returns"`
	// Schema where the function is located. If not specified, the provider default schema is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/function#schema Function#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

