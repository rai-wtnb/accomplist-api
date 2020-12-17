resource "aws_lb_listener" "https" {
  load_balancer_arn = aws_lb.lb.arn
  port              = "443"
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-2015-05"
  certificate_arn   = aws_acm_certificate.accomplist-cert.arn
  default_action {
    target_group_arn = aws_lb_target_group.http.arn
    type             = "forward"
  }
}
