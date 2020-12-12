resource "aws_route53_record" "accomplist_certificate" {
  zone_id = data.aws_route53_zone.accomplist_zone.id
  name    = tolist(aws_acm_certificate.accomplist-cert.domain_validation_options)[0].resource_record_name
  type    = tolist(aws_acm_certificate.accomplist-cert.domain_validation_options)[0].resource_record_type
  records = [
    tolist(aws_acm_certificate.accomplist-cert.domain_validation_options)[0].resource_record_value
  ]
  ttl = 60
}
