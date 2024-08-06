# Running a Tofu with the pokus provider

## How to use the Pokus Provider


* Create or update all by running:

```Powershell
$env:TF_LOG = "debug"

tofu init

tofu validate
tofu fmt

tofu plan
tofu apply -auto-approve

```

* Delete all by running:

```Powershell

# destroy it all, and recreate it all:

tofu plan -destroy -out="my.first.destroy.plan.tfplan"; tofu apply "my.first.destroy.plan.tfplan";

```

* Delete all and recreate everything from scratch:

```Powershell

# destroy it all, and recreate it all:

tofu plan -destroy -out="my.first.destroy.plan.tfplan"; tofu apply "my.first.destroy.plan.tfplan"; tofu plan -out="my.first.plan.tfplan"; tofu apply -auto-approve "my.first.plan.tfplan"

```

## How this Tofu Provider Project was spinned up

### Prepare the source code

* In the `./examples/terraform/provider/example1/terraform-provider-pokus` folder, we have git cloned the code of the provider template https://github.com/hashicorp/terraform-provider-scaffolding-framework, and executed :

```bash
export TF_SRC_CODE_FOLDER="./examples/terraform/provider/example1/terraform-provider-pokus"

git clone https://github.com/hashicorp/terraform-provider-scaffolding-framework ${TF_SRC_CODE_FOLDER}


cd ${TF_SRC_CODE_FOLDER}

go mod edit -module terraform-provider-pokus

go mod tidy

go get github.com/3forges/pesto-api-client-go@v0.0.11

sed -i "s#github.com/hashicorp/terraform-provider-scaffolding-framework/internal/provider#terraform-provider-pokus/internal/provider#g" ./main.go


```

* Then change the `./internal/provider/provider.go` content so that it is:

```Golang
package provider

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
    "github.com/hashicorp/terraform-plugin-framework/provider"
    "github.com/hashicorp/terraform-plugin-framework/provider/schema"
    "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
    _ provider.Provider = &pokusProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
    return func() provider.Provider {
        return &pokusProvider{
            version: version,
        }
    }
}

// pokusProvider is the provider implementation.
type pokusProvider struct {
    // version is set to the provider version on release, "dev" when the
    // provider is built and ran locally, and "test" when running acceptance
    // testing.
    version string
}

// Metadata returns the provider type name.
func (p *pokusProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
    resp.TypeName = "pokus"
    resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *pokusProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
    resp.Schema = schema.Schema{}
}

// Configure prepares a HashiCups API client for data sources and resources.
func (p *pokusProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

// DataSources defines the data sources implemented in the provider.
func (p *pokusProvider) DataSources(_ context.Context) []func() datasource.DataSource {
    return nil
}

// Resources defines the resources implemented in the provider.
func (p *pokusProvider) Resources(_ context.Context) []func() resource.Resource {
    return nil
}

```

* Then change the `main.go` content so that it is:

```Golang
func main() {
    var debug bool

    flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
    flag.Parse()

    opts := providerserver.ServeOpts{
        // NOTE: This is not a typical Terraform Registry provider address,
        // such as registry.terraform.io/hashicorp/pokus. This specific
        // provider address is used in these tutorials in conjunction with a
        // specific Terraform CLI configuration for manual development testing
        // of this provider.
        Address: "pokus-io.io/terraform/pokus",
        Debug:   debug,
    }

    err := providerserver.Serve(context.Background(), provider.New(version), opts)

    if err != nil {
        log.Fatal(err.Error())
    }
}

```

### Prepare the test

Cd into the `./examples/terraform/provider/example1/` folder of this repo.

#### Windows

##### Step 1: Prepare the `terraform.rc`

* In git bash for windows, run:

```bash
export TF_RC_FILE_PATH=$(echo $APPDATA | sed 's#\\#/#g' | sed 's#C:#/c#g')/terraform.rc
echo "  TF_RC_FILE_PATH=[${TF_RC_FILE_PATH}]"

export MY_GO_BIN_FOLDER=$(go env GOBIN)
export MY_DEFAULT_GO_BIN_FOLDER=$(echo "$HOME/go/bin" | sed 's#/c##')

if [ "x${MY_GO_BIN_FOLDER}" == "x" ]; then
  echo "GO BIN is empty, so will use default"
  export PATH_FOR_DEV_OVERRIDES="${MY_DEFAULT_GO_BIN_FOLDER}"
else
  echo "GO BIN is not empty"
  export PATH_FOR_DEV_OVERRIDES="${MY_DEFAULT_GO_BIN_FOLDER}"
fi;


echo "  PATH_FOR_DEV_OVERRIDES=[${PATH_FOR_DEV_OVERRIDES}]"


cat <<EOF >${TF_RC_FILE_PATH}
provider_installation {

  dev_overrides {
      # "pokus-io.io/terraform/pokus" = "<PATH>"
      "pokus-io.io/terraform/pokus" = "${PATH_FOR_DEV_OVERRIDES}"
      }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}

EOF
```

Above, the `pokus-io.io/terraform/pokus` must match the value set in this block, in the `main.go` of your provider:

```Golang
    opts := providerserver.ServeOpts{
        // NOTE: This is not a typical Terraform Registry provider address,
        // such as registry.terraform.io/hashicorp/hashicups. This specific
        // provider address is used in these tutorials in conjunction with a
        // specific Terraform CLI configuration for manual development testing
        // of this provider.
        Address: "pokus-io.io/terraform/pokus",
        Debug:   debug,
    }
```

##### Step 2: go install your provider

To make our developed terraform provider, available in the GOBIN folder, go to the `./examples/terraform/provider/example1/terraform-provider-pokus` folder, and run, either in powershell or git bash for windows :

```bash
go install .
```

##### Step 3: run a TOFU/terraform recipe using the pokus provider

Now, we are ready to start running TOFU/terraform recipe.

Cd into the `./examples/terraform/provider/example1/` folder of this repo, and create a `main.tf` file with the below content:

```Terraform
terraform {
  required_providers {
    pokus = {
      source = "pokus-io.io/terraform/pokus"
    }
  }
}

provider "pokus" {}

data "pokus_coffees" "example" {}
```

And in powershell, in the same folder, run:

```Powershell
$env:TF_LOG = "debug"

tofu init

tofu validate
tofu fmt

tofu plan
tofu apply -auto-approve

```

What tells us that the setup is successful, is that the DEBUG logs will confirm the pokus terraform plugin is running, and an error is thrown telling that the `pokus_coffees` datasource does not exist, which makes sense, as far as of now, we did not define any datasource of name `pokus_coffees` in our provider code inside the `./examples/terraform/provider/example1/terraform-provider-pokus` folder:

```bash
2024-08-03T13:09:43.313+0200 [DEBUG] provider: using plugin: version=6
2024-08-03T13:09:43.373+0200 [ERROR] vertex "data.pokus_coffees.example" error: Invalid data source
╷
│ Error: Invalid data source
│
│   on main.tf line 11, in data "pokus_coffees" "example":
│   11: data "pokus_coffees" "example" {}
│
│ The provider pokus-io.io/terraform/pokus does not support data source "pokus_coffees".
╵
2024-08-03T13:09:43.381+0200 [DEBUG] provider.stdio: received EOF, stopping recv loop: err="rpc error: code = Unavailable desc = error reading from server: EOF"
```
