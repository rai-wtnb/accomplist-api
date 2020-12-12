resource "aws_acm_certificate" "accomplist-cert" {
  domain_name               = data.aws_route53_zone.accomplist_zone.name
  subject_alternative_names = []
  validation_method         = "DNS"
}
