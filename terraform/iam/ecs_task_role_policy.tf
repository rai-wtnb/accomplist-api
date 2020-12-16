resource "aws_iam_policy" "ecs_task_role_policy" {
  name = "ecs_task_role_policy"
  path = "/"
  description = "ecs task role policy"
  policy = file("policies/ecs_task_policy.json")
  }
