resource "aws_instance" "accomplist" {
  ami                         = "ami-0e37e42dff65024ae"
  instance_type               = "t2.small"
  key_name                    = aws_key_pair.accomplist_ec2_key_pair.key_name
  monitoring                  = true
  iam_instance_profile        = data.terraform_remote_state.aws_iam.outputs.ecs_instance_profile_name
  subnet_id                   = data.terraform_remote_state.vpc.outputs.public_subnet_1_id
  user_data                   = file("./user_data.sh")
  associate_public_ip_address = true
  vpc_security_group_ids = [
    "${aws_security_group.instance.id}",
  ]
  root_block_device {
    volume_size = "30"
    volume_type = "gp2"
  }

  tags = {
    "Name" = "accomplist-ec2"
  }
}

resource "aws_key_pair" "accomplist_ec2_key_pair" {
  key_name = var.aws_key_name
  public_key = file(var.public_key_path)

  tags = {
    Name  = "accomplist-instance-key-pair"
  }
}
