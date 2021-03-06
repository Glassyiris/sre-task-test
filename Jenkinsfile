#!groovy
def APP_NAME='entry-task-weimin'
def GROUP='shopee' // harbor.shopeemobile.com/$GROUP/image:tag
def POD_PORT='8080'
def createVersion() {
    return new Date().format('yyyyMMddHHmmss')
}
def BuildGitCommit = createVersion()

node ('slave') {
 properties([
 parameters([
 string(name: 'NAMESPACE', defaultValue: 'weimin', description: ''),
 string(name: 'REPLICAS', defaultValue: '1', description: ''),
 string(name: 'CPU', defaultValue: '1', description: ''),
 string(name: 'MEMORY', defaultValue: '1', description: ''),
 ])
 ])

 def myRepo = checkout scm
 def gitCommit = myRepo.GIT_COMMIT
 def gitBranch = myRepo.GIT_BRANCH
 def shortGitCommit = "${BuildGitCommit}"
 def previousGitCommit = sh(script: "git rev-parse ${gitCommit}~", returnStdout: true)


 def REGISTRY_CRED = 'harbor-szdevops'
 def BUILD_TAG = "git-${shortGitCommit}"
 def IMAGE_NAME = "harbor.shopeemobile.com/${GROUP}/${APP_NAME}:${BUILD_TAG}"

 stage('Build') {
 def customImage = docker.build("${IMAGE_NAME}", "-f Dockerfile .")
 docker.withRegistry('https://harbor.shopeemobile.com', "${REGISTRY_CRED}") {
 customImage.push()
 }
 }

 stage('Config') {
 sh """
 sed -i '
 s#\$IMAGE_NAME#$IMAGE_NAME#g
 s#\$NAMESPACE#$NAMESPACE#g
 s#\$REPLICAS#$REPLICAS#g
 s#\$APP_NAME#$APP_NAME#g
 s#\$POD_PORT#$POD_PORT#g
 s#\$CPU#$CPU#g
 s#\$MEMORY#${MEMORY}Gi#g
 ' deploy.yml
 """
 }

 stage('Deploy'){
 def kubeConfigId = "credit-k8s-test-sg2"

 kubernetesDeploy(kubeconfigId: "${kubeConfigId}",
 configs: 'deploy.yml',
 enableConfigSubstitution: true,
 )
 }
}
