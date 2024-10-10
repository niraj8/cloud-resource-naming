# cloud-resource-naming

## AWS

### naming validations

s3: https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html
- bucket name rules
- object name rules: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html
- directory bucket name rules: https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-bucket-naming-rules.html

dynamodb: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html
- table name rules
- attribute name rules
- index name rules
  
rds
alb target group
elasticache group
ec2 instance type: https://docs.aws.amazon.com/ec2/latest/instancetypes/instance-type-names.html
- comments in this thread are interesting: https://stackoverflow.com/questions/46052869/what-are-the-most-restrictive-aws-resource-name-limitations-e-g-characters-and

arn validation

cloudtrail
- s3 bucket you create for cloudtrail has special requirements
https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-trail-naming-requirements.html
https://medium.com/cloud-security/aws-resources-organization-and-naming-conventions-262676d6e202
## naming conventions


later
- https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-label-requirements.html
