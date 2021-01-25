provider "azurerm" {
  version = "2.22.0"
  features {}
}

provider "azuread" {
  version = "~> 0.11"
  features {}
}

provider "aws" {
  version = "2.70.0"
}


provider "helm" {
  version = "1.2.3"
}


resource "azurerm_resource_group" "example" {
  name     = "example"
  location = "West Europe"
}