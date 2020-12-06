resource "aws_iam_policy" "deploy" {
  name = "deploy"
  path = "/"
  description = "deploy policy"
  policy = file("policies/ecr_policy.json")
  }
