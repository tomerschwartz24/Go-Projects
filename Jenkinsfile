 
pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                echo 'Checking out code from repository'
                checkout scm
            }
        }

        stage('Build') {
            steps {
                sh 'docker build -t tomerschwartz2411/website:webapp -f Dockerfile .'
            }
        }
        

        // Staging is jenkins server localhost due to the fact that this is a homelab.
        stage('Test') {
            steps {
                script {
                sh "docker run  --name staging_$BUILD_NUMBER  -p 80:80 -d tomerschwartz2411/website:webapp"
                // curl test to check if application is accessible via http.
                def curlExitCode = sh(script: "curl -s localhost:80 |grep -i 'ts devops enginner'", returnStatus: true)
                if (curlExitCode != 0) {
                    sh "sudo docker stop staging_$BUILD_NUMBER  > /dev/null 2>&1  || true"
                    // If the application fails the curl test the pipeline should fail at this condition.
                    error "Unable to determine if website is working properly, exit code : ${curlExitCode}"
                } else {
                    println "Website seems to be functional, continuing..."}
                    sh "sudo docker stop staging_$BUILD_NUMBER  > /dev/null 2>&1  || true"
                }
            }

        }

        stage('Push') {
            steps {
                withCredentials([usernamePassword( credentialsId: 'docker_hub', usernameVariable: 'USER', passwordVariable: 'PASSWORD')]) {
                sh 'docker login -u $USER -p $PASSWORD docker.io'
                sh 'docker push tomerschwartz2411/website:webapp'
                } 
            }
        }
        stage('Deploy to prod') {
            steps {
            script {
                //Running WatchTower enables continuous deployment, the application images will be updated with  minimal downtime, 
                //combined with Jenkins poll SCM or Webhook this will allow a fully automated process.
                //WatchTower is searching for an updated digest of the currently used image name in the docker hub registry, if there is one it will deploy it instead of the existing one.
              sh '''
                 ssh -i /opt/ssh-access.pem -o StrictHostKeyChecking=no ubuntu@tomerschwartz.com \
                 sudo  docker run  -e REPO_USER=$USER -e REPO_PASS=$PASSWORD -v /var/run/docker.sock:/var/run/docker.sock containrrr/watchtower --run-once --debug
                '''
                }
            }
        }
    } 
}

