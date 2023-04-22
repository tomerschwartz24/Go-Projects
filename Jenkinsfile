 
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
                    error "Unable to determine if website is working properly, exit code : ${curlExitCode}"
                } else {
                    println "Website seems to be functional, continuing..."}
                }
            }
        }
    } 
}

