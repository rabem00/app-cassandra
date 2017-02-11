# app-cassandra
A test Go app using Apache Cassandra


# Install Cassandra on CentOS 7

yum -y update
yum -y install java
vi /etc/yum.repos.d/datastax.repo
  [datastax]
  name = DataStax Repo for Apache Cassandra
  baseurl = http://rpm.datastax.com/community
  enabled = 1
  gpgcheck = 0

yum -y install dsc20
systemctl enable cassandra.service

# Add ports on internal interface for Cassandra server
firewall-cmd --zone=internal --add-port=7000/tcp --add-port=7199/tcp --add-port=9042/tcp --add-port=9160/tcp --add-port=61619-61621/tcp --permanent
# Add ports on public interface for Cassandra server
firewall-cmd --zone=public --add-port=80/tcp --add-port=8888/tcp --permanent
firewall-cmd --reload

# Test connection and connect with cqlsh
nodetool status
cqlsh
