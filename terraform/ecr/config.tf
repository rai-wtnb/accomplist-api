variable "prefix" {
  default = "accomplist"
}

terraform {
  backend "s3" {
    bucket = "accomplist-tfstate-bucket"
    key    = "accomplist/ecr/terraform.tfstate"
    region = "ap-northeast-1"
  }
}

provider "aws" {
  region = "ap-northeast-1"
}
