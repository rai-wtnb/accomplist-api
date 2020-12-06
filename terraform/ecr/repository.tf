resource "aws_ecr_repository" "accomplist-ecr" {
  name = "${var.prefix}-ecr"
}
