resource "aws_ecs_service" "accomplist-api" {
  name            = "accomplist-api"
  cluster         = aws_ecs_cluster.accomplist-ecs-cluster.id
  task_definition = aws_ecs_task_definition.accomplist-task.arn
  desired_count   = 1
  launch_type     = "EC2"
}
