resource "aws_s3_bucket" "bucket" {
  bucket = "${var.prefix}-bucket"
  acl    = "private"

  tags = {
    Name        = "${var.prefix}-bucket"
    Environment = "common"
  }
}
