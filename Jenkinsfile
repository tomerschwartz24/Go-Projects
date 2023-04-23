 
pipeline {
    agent any
    stages {
        stage('Building Website') {
            steps {
                sh 'docker build -t tomerschwartz2411/website:webapp_$BUILD_NUMBER -f Dockerfile .'
            }
        }
        
        stage('Pushing to DockerHub') {
            steps {
                withCredentials([usernamePassword( credentialsId: 'docker_hub', usernameVariable: 'USER', passwordVariable: 'PASSWORD')]) {
                sh 'docker login -u $USER -p $PASSWORD docker.io'
                sh 'docker push tomerschwartz2411/website:webapp_$BUILD_NUMBER'
                } 
            }
        }
        // Staging is jenkins server localhost due to the fact that this is a homelab.
        stage('Deploy to Staging & Test') {
            steps {
                script {
                sh "docker run  --name staging_$BUILD_NUMBER -p 80:80 -d tomerschwartz2411/website:webapp_$BUILD_NUMBER"
                def curlExitCode = sh(script: "curl -s localhost:80 |grep -i 'ts devops enginner'", returnStatus: true)
                if (curlExitCode != 0) {
                    sh "sudo docker stop $(sudo docker ps -q) > /dev/null 2>&1  || true"
                    error "Unable to determine if website is working properly, exit code : ${curlExitCode}"
                } else {
                    println "Website seems to be functional, continuing..."}
                    sh "sudo docker stop $(sudo docker ps -q) > /dev/null 2>&1  || true"
                }
            }

        }
        stage('Deploy to Production') {
            steps {
            script {
              def current_runtime_tag_webapp = "tomerschwartz2411/website:webapp_$BUILD_NUMBER"
              
              sh '''
                 ssh -i /opt/ssh-access.pem -o StrictHostKeyChecking=no ubuntu@tomerschwartz.com \
                  sudo docker pull ''' +current_runtime_tag_webapp+ ''' && \
                  sudo docker stop $(sudo docker ps -q) > /dev/null 2>&1  || true && \
                  sudo docker run -p 80:80 -d ''' +current_runtime_tag_webapp+ ''' 
                '''
                }
            }
        }
    } 
}

