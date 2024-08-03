---
title: Chapter 1
# slug: my-custom-slug/supports/slashes
tags:
  - golang
  - basics

summary: How to implement a TOFU (terraform) provider
index: 0
---

## First try: the official terraform tutorial

In tis chapter, I will try and implement a tofu provider, using hashcorp's official <https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider>

* There:

```bash
git clone https://github.com/hashicorp/terraform-provider-scaffolding-framework ./terraform-provider-pokus

cd ./terraform-provider-pokus/

# ---
# then rename the go module
# go mod edit -module terraform-provider-hashicups
go mod edit -module terraform-provider-pokus

# ---
# Then, install all the provider's dependencies.
go mod tidy

```

Note there that at the date I write this, 03/08/2024, the <https://github.com/hashicorp/terraform-provider-scaffolding-framework> has no change of licence, like the new business license for hashicorp. It has a mozilla public license.


### Step 1: reset internal dependencies

Open the `main.go` file in the `terraform-provider-pokus` repository's root directory and replace the import declaration that you find :

```Golang
import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-provider-scaffolding-framework/internal/provider"
)
```

with the following:

```Golang
import (
    "context"
    "flag"
    "log"

    "github.com/hashicorp/terraform-plugin-framework/providerserver"

    "terraform-provider-pokus/internal/provider"
)

```

### Step 2: get a rest api up and running

I will try with my own api available at :

* <http://api.pesto.io:3000/api>
* and as graphql <http://api.pesto.io:3000/graphql>

if I experience any issue, i can always use the "hashicups" rest api from the tutorial.



### Step 3: The most simple `provider.go` and `main.go`

Edit the `provider.go` and `main.go` files, with the below content:


* In `provider.go`:

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

* In `main.go`

```Golang
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
    "context"
    "flag"
    "log"

    "github.com/hashicorp/terraform-plugin-framework/providerserver"

    "terraform-provider-pokus/internal/provider"
)


// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate -provider-name scaffolding

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	version string = "dev"

	// goreleaser can pass other information to the main package, such as the specific commit
	// https://goreleaser.com/cookbooks/using-main.version/
)
func main() {
    var debug bool

    flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
    flag.Parse()

    opts := providerserver.ServeOpts{
        // NOTE: This is not a typical Terraform Registry provider address,
        // such as registry.terraform.io/hashicorp/hashicups. This specific
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

### Step 4: Prepare the `Tofu` dev env


Cd into the `./examples/terraform/provider/example1/` folder of this repo.

#### Windows

##### Point 1: Prepare the `terraform.rc`

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

##### Point 2: go install your provider

To make our developed terraform provider, in the GOBIN folder, go to the `vvvv` folder, and run, either in powershell or git bash for windows :

```bash
go install .
```

##### Point 3: run a TOFU/terraform recipe using the pokus provider

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

ANd in powershell, in the same folder, run:

```Powershell
$env:TF_LOG = "debug"

tofu init

tofu validate
tofu fmt

tofu plan

```

What tells us that the setup is successful, is that the DEBUG logs will confirm the pokus terraform plugin is running, and an error is thrown telling that the `pokus_coffees` datasource does not exist, which makes sense, as far as of now, we did not define any datasource of name `pokus_coffees` in our code:

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

### Step 5: Configure the client

The client is what is going to connect to my REST API.
It will hold the hostname, port num, URL to the REST API, and if necessary the auth credentials.

ALl the code of that client will be in the `provider.go` source code file.

The client's design is simple:

> The provider uses the `Configure` method to read API client configuration values from the Terraform configuration or environment variables. After verifying the values should be acceptable, the API client is created and made available for data source and resource usage.

The configure method follows these steps:

* **Retrieves values from the configuration**. The method will attempt to retrieve values from the provider configuration and convert it to an providerModel struct.
* **Checks for unknown configuration values**. The method prevents an unexpectedly misconfigured client, if Terraform configuration values are only known after another resource is applied.
* **Retrieves values from environment variables**. The method retrieves values from environment variables, then overrides them with any set Terraform configuration values.
* **Creates API client**. The method invokes the HashiCups API client's NewClient function.
* **Stores configured client for data source and resource usage**. The method sets the DataSourceData and ResourceData fields of the response, so the client is available for usage by data source and resource implementations.

The code presented by the hashicorp tutorial suggests to use an external go client able to hit the targetted rest API.

I created a go client for the pesto-api, <https://github.com:3forges/pesto-api-client-go> (fully inspired by the <https://github.com/hashicorp-demoapp/hashicups-client-go>), and will use that Go client in the implementation of my teraform pokus provider, which will be able to CRUD Pesto API _project_ resources.

And I run in the `./examples/terraform/provider/example1/terraform-provider-pokus` folder, I run :

```bash
go get github.com/3forges/pesto-api-client-go

```

Then I edit my `provider.go` to add the **`Configure`** function (and VERY IMPORTANT, the schame function needs to be updated as well, addtionnally to the new `pokusProviderModel` `struct`) :

```Golang

// Schema defines the provider-level schema for configuration data.
func (p *pokusProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional: true,
			},
			"username": schema.StringAttribute{
				Optional: true,
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}



//////////////////////////////////////////////
//////////////////////////////////////////////
//////////////////////////////////////////////
//// BEGIN CONFIGURE PART
//////////////////////////////////////////////
//////////////////////////////////////////////
//////////////////////////////////////////////

// // Configure prepares a Pesto API client for data sources and resources.
// func (p *pokusProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
// }
// pokusProviderModel maps provider schema data to a Go type.
type pokusProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func (p *pokusProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
    // Retrieve provider data from configuration
    var config pokusProviderModel
    diags := req.Config.Get(ctx, &config)
    resp.Diagnostics.Append(diags...)
	
    if resp.Diagnostics.HasError() {
        return
    }

    // If practitioner provided a configuration value for any of the
    // attributes, it must be a known value.

    if config.Host.IsUnknown() {
        resp.Diagnostics.AddAttributeError(
            path.Root("host"),
            "Unknown HashiCups API Host",
            "The provider cannot create the HashiCups API client as there is an unknown configuration value for the HashiCups API host. "+
                "Either target apply the source of the value first, set the value statically in the configuration, or use the POKUS_HOST environment variable.",
        )
    }

    if config.Username.IsUnknown() {
        resp.Diagnostics.AddAttributeError(
            path.Root("username"),
            "Unknown HashiCups API Username",
            "The provider cannot create the HashiCups API client as there is an unknown configuration value for the HashiCups API username. "+
                "Either target apply the source of the value first, set the value statically in the configuration, or use the POKUS_USERNAME environment variable.",
        )
    }

    if config.Password.IsUnknown() {
        resp.Diagnostics.AddAttributeError(
            path.Root("password"),
            "Unknown HashiCups API Password",
            "The provider cannot create the HashiCups API client as there is an unknown configuration value for the HashiCups API password. "+
                "Either target apply the source of the value first, set the value statically in the configuration, or use the POKUS_PASSWORD environment variable.",
        )
    }

    if resp.Diagnostics.HasError() {
        return
    }

    // Default values to environment variables, but override
    // with Terraform configuration value if set.

    host := os.Getenv("POKUS_HOST")
    username := os.Getenv("POKUS_USERNAME")
    password := os.Getenv("POKUS_PASSWORD")

    if !config.Host.IsNull() {
        host = config.Host.ValueString()
    }

    if !config.Username.IsNull() {
        username = config.Username.ValueString()
    }

    if !config.Password.IsNull() {
        password = config.Password.ValueString()
    }

    // If any of the expected configurations are missing, return
    // errors with provider-specific guidance.

    if host == "" {
        resp.Diagnostics.AddAttributeError(
            path.Root("host"),
            "Missing HashiCups API Host",
            "The provider cannot create the HashiCups API client as there is a missing or empty value for the HashiCups API host. "+
                "Set the host value in the configuration or use the POKUS_HOST environment variable. "+
                "If either is already set, ensure the value is not empty.",
        )
    }

    if username == "" {
        resp.Diagnostics.AddAttributeError(
            path.Root("username"),
            "Missing HashiCups API Username",
            "The provider cannot create the HashiCups API client as there is a missing or empty value for the HashiCups API username. "+
                "Set the username value in the configuration or use the POKUS_USERNAME environment variable. "+
                "If either is already set, ensure the value is not empty.",
        )
    }

    if password == "" {
        resp.Diagnostics.AddAttributeError(
            path.Root("password"),
            "Missing HashiCups API Password",
            "The provider cannot create the HashiCups API client as there is a missing or empty value for the HashiCups API password. "+
                "Set the password value in the configuration or use the POKUS_PASSWORD environment variable. "+
                "If either is already set, ensure the value is not empty.",
        )
    }

    if resp.Diagnostics.HasError() {
        return
    }

    // Create a new PestoAPI client using the configuration values
    client, err := pesto.NewClient(&host, &username, &password)
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to Create HashiCups API Client",
            "An unexpected error occurred when creating the HashiCups API client. "+
                "If the error is not clear, please contact the provider developers.\n\n"+
                "HashiCups Client Error: "+err.Error(),
        )
        return
    }

    // Make the HashiCups client available during DataSource and Resource
    // type Configure methods.
    resp.DataSourceData = client
    resp.ResourceData = client
}

//////////////////////////////////////////////
//////////////////////////////////////////////
//////////////////////////////////////////////
//// END CONFIGURE PART
//////////////////////////////////////////////
//////////////////////////////////////////////
//////////////////////////////////////////////


```


TO complete with new datasource added

then run the tofu:

```Powershell
$env:POKUS_HOST = "http://api.pesto.io:3000"
$env:POKUS_USERNAME = "education"
$env:POKUS_PASSWORD = "test123"
tofu plan

```

```Powershell
export POKUS_HOST="http://api.pesto.io:3000"
export POKUS_USERNAME="education"
export POKUS_PASSWORD="test123"
tofu plan

```

Note that then, I get an error with the HTTP call to the pesto API: a 404 of course, because my API has no authentication endpoint.

I will have to fix that, one way, or another (is it possible to implement a terraform provider without authentication? )

Response: Yes! The authenticationwas made mandatory only in the go client of the REST API, with the signin method, I just commented it and it all worked, the GET Http call to the API successfully fetched the pesto-projects!

Actually not, the go client is just created: the datasource alone, even with a terraform output, does not fetch the api, it obviously is based only on the terraform state.
