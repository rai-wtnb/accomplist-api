variable "prefix" {
  default = "accomplist"
}
variable aws_key_name {}
variable public_key_path {}

terraform {
  backend "s3" {
    bucket = "accomplist-tfstate-bucket"
    key    = "accomplist/service/terraform.tfstate"
    region = "ap-northeast-1"
  }
}

provider "aws" {
  region = "ap-northeast-1"
}
