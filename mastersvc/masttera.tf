provider "aws" {

}
variable "AWS_REGION" {
  type = string
}
AWS_REGION="eu-west-1"

variable "AMIS" {
  type = map(string)
  default = {
    eu-west-1="my ami"
  }
}


resource "aws_instance" "masterservice" {
  ami = var.AMIS[var.AWS_REGION]
  instance_type = "t2.micro"
}