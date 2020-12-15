resource "aws_cloudwatch_log_group" "accomplist-api" {
  name              = "accomplist-api"
  retention_in_days = 5
}
