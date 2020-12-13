resource "aws_subnet" "public_subnet_1" {
  vpc_id            = "${aws_vpc.vpc.id}"
  cidr_block        = "10.0.0.0/20"
  availability_zone = "ap-northeast-1a"
  tags = {
    "Name" = "accomplist-public-subnet-1"
  }
}

resource "aws_subnet" "public_subnet_2" {
  vpc_id            = "${aws_vpc.vpc.id}"
  cidr_block        = "10.0.16.0/20"
  availability_zone = "ap-northeast-1c"
  tags = {
    "Name" = "accomplist-public-subnet-2"
  }
}

resource "aws_subnet" "private_subnet_1" {
  vpc_id            = "${aws_vpc.vpc.id}"
  cidr_block        = "10.0.32.0/20"
  availability_zone = "ap-northeast-1a"
  tags = {
    "Name" = "accomplist-private-subnet-1"
  }
}

resource "aws_subnet" "private_subnet_2" {
  vpc_id            = "${aws_vpc.vpc.id}"
  cidr_block        = "10.0.48.0/20"
  availability_zone = "ap-northeast-1c"
  tags = {
    "Name" = "accomplist-private-subnet-2"
  }
}
