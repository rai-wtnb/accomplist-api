data "aws_route53_zone" "accomplist_zone" {
  name = "accomplist-api.com"
}

resource "aws_route53_record" "accomplist-record" {
  zone_id = data.aws_route53_zone.accomplist_zone.zone_id
  name    = data.aws_route53_zone.accomplist_zone.name
  type    = "A"

  alias {
    name                   = aws_lb.lb.dns_name
    zone_id                = aws_lb.lb.zone_id
    evaluate_target_health = true
  }
}
