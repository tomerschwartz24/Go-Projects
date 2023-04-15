pipeline {
    agent any
    stages {
        stage('Building website binary') {
            steps {
              sh 'go build website.go'
            }
        }
        stage('Building Website') {
            steps {
                sh 'docker build -t tomerschwartz2411/website:webapp -f Dockerfile .'
            }
        }
        
        stage('Pushing to DockerHub') {
            steps {
                withCredentials([usernamePassword( credentialsId: 'docker_hub', usernameVariable: 'USER', passwordVariable: 'PASSWORD')]) {
                sh 'docker login -u $USER -p $PASSWORD docker.io'
                sh 'docker push tomerschwartz2411/website:webapp'
                }
            }
        }
    }
}