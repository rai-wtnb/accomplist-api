resource "aws_lb_target_group" "http" {
  name     = "accomplist-http"
  port     = 80
  protocol = "HTTP"
  vpc_id   = data.terraform_remote_state.vpc.outputs.vpc_id
}
