connection "jenkins" {
  plugin = "jenkins"

  # The Jenkins server URL is required for all requests. Required.
  # It should be fully qualified (e.g. # https://...) and point to the root of the Jenkins server location.
  # Can also be set via the JENKINS_URL environment variable.
  # server_url = "https://ci-cd.internal.my-company.com"

  # The Jenkins username for authentication is required for requests. Required.
  # Can also be set via the JENKINS_USERNAME environment variable.
  # username = "admin"

  # Either the Jenkins password or the API token is required for requests. Required. 
  # Can also be set via the JENKINS_PASSWORD environment variable.
  # password = "aqt*abc8vcf9abc.ABC"

  # Further information: https://www.jenkins.io/doc/book/using/using-credentials/   
}
