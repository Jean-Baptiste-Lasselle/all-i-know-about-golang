# Running a Tofu with the pokus provider

## Prepare the dev env

Cd into the `./examples/terraform/provider/example1/` folder of this repo.

### Windows

#### Step 1: Prepare the `terraform.rc`

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


#### Step 2: go install your provider

To make our developed terraform provider, in the GOBIN folder, go to the `vvvv` folder, and run, either in powershell or git bash for windows :

```bash
go install .
```

#### Step 3: run a TOFU/terraform recipe using the pokus provider

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
